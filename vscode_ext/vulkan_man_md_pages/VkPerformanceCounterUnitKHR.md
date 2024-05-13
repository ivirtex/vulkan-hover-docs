# VkPerformanceCounterUnitKHR(3) Manual Page

## Name

VkPerformanceCounterUnitKHR - Supported counter unit types



## <a href="#_c_specification" class="anchor"></a>C Specification

Performance counters have an associated unit. This unit describes how to
interpret the performance counter result.

The performance counter unit types which **may** be returned in
[VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterKHR.html)::`unit` are:

``` c
// Provided by VK_KHR_performance_query
typedef enum VkPerformanceCounterUnitKHR {
    VK_PERFORMANCE_COUNTER_UNIT_GENERIC_KHR = 0,
    VK_PERFORMANCE_COUNTER_UNIT_PERCENTAGE_KHR = 1,
    VK_PERFORMANCE_COUNTER_UNIT_NANOSECONDS_KHR = 2,
    VK_PERFORMANCE_COUNTER_UNIT_BYTES_KHR = 3,
    VK_PERFORMANCE_COUNTER_UNIT_BYTES_PER_SECOND_KHR = 4,
    VK_PERFORMANCE_COUNTER_UNIT_KELVIN_KHR = 5,
    VK_PERFORMANCE_COUNTER_UNIT_WATTS_KHR = 6,
    VK_PERFORMANCE_COUNTER_UNIT_VOLTS_KHR = 7,
    VK_PERFORMANCE_COUNTER_UNIT_AMPS_KHR = 8,
    VK_PERFORMANCE_COUNTER_UNIT_HERTZ_KHR = 9,
    VK_PERFORMANCE_COUNTER_UNIT_CYCLES_KHR = 10,
} VkPerformanceCounterUnitKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PERFORMANCE_COUNTER_UNIT_GENERIC_KHR` - the performance counter
  unit is a generic data point.

- `VK_PERFORMANCE_COUNTER_UNIT_PERCENTAGE_KHR` - the performance counter
  unit is a percentage (%).

- `VK_PERFORMANCE_COUNTER_UNIT_NANOSECONDS_KHR` - the performance
  counter unit is a value of nanoseconds (ns).

- `VK_PERFORMANCE_COUNTER_UNIT_BYTES_KHR` - the performance counter unit
  is a value of bytes.

- `VK_PERFORMANCE_COUNTER_UNIT_BYTES_PER_SECOND_KHR` - the performance
  counter unit is a value of bytes/s.

- `VK_PERFORMANCE_COUNTER_UNIT_KELVIN_KHR` - the performance counter
  unit is a temperature reported in Kelvin.

- `VK_PERFORMANCE_COUNTER_UNIT_WATTS_KHR` - the performance counter unit
  is a value of watts (W).

- `VK_PERFORMANCE_COUNTER_UNIT_VOLTS_KHR` - the performance counter unit
  is a value of volts (V).

- `VK_PERFORMANCE_COUNTER_UNIT_AMPS_KHR` - the performance counter unit
  is a value of amps (A).

- `VK_PERFORMANCE_COUNTER_UNIT_HERTZ_KHR` - the performance counter unit
  is a value of hertz (Hz).

- `VK_PERFORMANCE_COUNTER_UNIT_CYCLES_KHR` - the performance counter
  unit is a value of cycles.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_performance_query.html),
[VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceCounterUnitKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
