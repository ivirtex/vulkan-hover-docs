# VkPipelineExecutableStatisticValueKHR(3) Manual Page

## Name

VkPipelineExecutableStatisticValueKHR - A union describing a pipeline executable statistic



## [](#_c_specification)C Specification

The `VkPipelineExecutableStatisticValueKHR` union is defined as:

```c++
// Provided by VK_KHR_pipeline_executable_properties
typedef union VkPipelineExecutableStatisticValueKHR {
    VkBool32    b32;
    int64_t     i64;
    uint64_t    u64;
    double      f64;
} VkPipelineExecutableStatisticValueKHR;
```

## [](#_members)Members

- `b32` is the 32-bit boolean value if the `VkPipelineExecutableStatisticFormatKHR` is `VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_BOOL32_KHR`.
- `i64` is the signed 64-bit integer value if the `VkPipelineExecutableStatisticFormatKHR` is `VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_INT64_KHR`.
- `u64` is the unsigned 64-bit integer value if the `VkPipelineExecutableStatisticFormatKHR` is `VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_UINT64_KHR`.
- `f64` is the 64-bit floating-point value if the `VkPipelineExecutableStatisticFormatKHR` is `VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_FLOAT64_KHR`.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_executable\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_executable_properties.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkPipelineExecutableStatisticKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutableStatisticKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineExecutableStatisticValueKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0