# VkPipelineExecutableStatisticFormatKHR(3) Manual Page

## Name

VkPipelineExecutableStatisticFormatKHR - Enum describing a pipeline executable statistic



## [](#_c_specification)C Specification

The `VkPipelineExecutableStatisticFormatKHR` enum is defined as:

```c++
// Provided by VK_KHR_pipeline_executable_properties
typedef enum VkPipelineExecutableStatisticFormatKHR {
    VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_BOOL32_KHR = 0,
    VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_INT64_KHR = 1,
    VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_UINT64_KHR = 2,
    VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_FLOAT64_KHR = 3,
} VkPipelineExecutableStatisticFormatKHR;
```

## [](#_description)Description

- `VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_BOOL32_KHR` specifies that the statistic is returned as a 32-bit boolean value which **must** be either `VK_TRUE` or `VK_FALSE` and **should** be read from the `b32` field of `VkPipelineExecutableStatisticValueKHR`.
- `VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_INT64_KHR` specifies that the statistic is returned as a signed 64-bit integer and **should** be read from the `i64` field of `VkPipelineExecutableStatisticValueKHR`.
- `VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_UINT64_KHR` specifies that the statistic is returned as an unsigned 64-bit integer and **should** be read from the `u64` field of `VkPipelineExecutableStatisticValueKHR`.
- `VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_FLOAT64_KHR` specifies that the statistic is returned as a 64-bit floating-point value and **should** be read from the `f64` field of `VkPipelineExecutableStatisticValueKHR`.

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_executable\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_executable_properties.html), [VkPipelineExecutableStatisticKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutableStatisticKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineExecutableStatisticFormatKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0