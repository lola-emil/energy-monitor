

option = {
  backgroundColor: 'transparent',

    tooltip: {
    trigger: 'axis',
      backgroundColor: "oklch(0.21 0.006 285.885)",
        borderColor: "oklch(1 0 0 / 10%)",
          textStyle: { color: "oklch(0.985 0 0)" }
  },

  grid: {
    left: '3%',
      right: '3%',
        top: 20,
          bottom: 40
  },

  xAxis: {
    type: 'category',
      data: [
        'Apr 1', 'Apr 2', 'Apr 3', 'Apr 4', 'Apr 5', 'Apr 6',
        'Apr 7', 'Apr 8', 'Apr 9', 'Apr 10', 'Apr 11', 'Apr 12',
        'Apr 13', 'Apr 14', 'Apr 15', 'Apr 16', 'Apr 17', 'Apr 18',
        'Apr 19', 'Apr 20', 'Apr 21', 'Apr 22', 'Apr 23', 'Apr 24',
        'Apr 25', 'Apr 26', 'Apr 27', 'Apr 28', 'Apr 29', 'Apr 30',
        'May 1', 'May 2', 'May 3', 'May 4', 'May 5', 'May 6',
        'May 7', 'May 8', 'May 9', 'May 10', 'May 11', 'May 12',
        'May 13', 'May 14', 'May 15', 'May 16', 'May 17', 'May 18',
        'May 19', 'May 20', 'May 21', 'May 22', 'May 23', 'May 24',
        'May 25', 'May 26', 'May 27', 'May 28', 'May 29', 'May 30',
        'May 31', 'Jun 1', 'Jun 2', 'Jun 3', 'Jun 4', 'Jun 5',
        'Jun 6', 'Jun 7', 'Jun 8', 'Jun 9', 'Jun 10', 'Jun 11',
        'Jun 12', 'Jun 13', 'Jun 14', 'Jun 15', 'Jun 16', 'Jun 17',
        'Jun 18', 'Jun 19', 'Jun 20', 'Jun 21', 'Jun 22', 'Jun 23',
        'Jun 24', 'Jun 25', 'Jun 26', 'Jun 27', 'Jun 28', 'Jun 29'
      ],
        axisLine: { lineStyle: { color: "oklch(1 0 0 / 10%)" } },
    axisLabel: {
      color: "oklch(0.274 0.006 286.033)",
        interval: 12 // show fewer labels like screenshot
    }
  },

  yAxis: {
    type: 'value',
      axisLabel: { color: "oklch(0.274 0.006 286.033)" },
    splitLine: {
      lineStyle: {
        color: "oklch(1 0 0 / 10%)",
          opacity: 0.3
      }
    }
  },

  series: [
    {
      type: 'bar',
      data: [
        137, 235, 244, 79, 166, 230, 119, 186, 112, 141, 72, 81,
        181, 238, 71, 197, 70, 199, 86, 110, 94, 96, 117, 226,
        81, 215, 93, 202, 112, 136, 138, 179, 95, 225, 245, 58,
        130, 139, 76, 179, 109, 69, 96, 176, 203, 235, 194, 79,
        218, 112, 202, 111, 113, 148, 227, 180, 213, 235, 190, 140,
        202, 228, 140, 221, 196, 196, 141, 225, 203, 164, 53, 237,
        237, 187, 187, 167, 175, 84, 180, 165, 149, 174, 57, 154,
        205, 222, 85, 116, 233, 72
      ],
      barWidth: 6, // thin bars
      itemStyle: {
        borderRadius: [2, 2, 0, 0],
        color: "oklch(0.637 0.237 25.331)"
      },
      emphasis: {
        itemStyle: {
          color: "oklch(0.637 0.237 25.331)" // slight variation on hover
        }
      }
    }
  ]
};