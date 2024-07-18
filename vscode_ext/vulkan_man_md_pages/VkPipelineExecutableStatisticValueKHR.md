# VkPipelineExecutableStatisticValueKHR(3) Manual Page

## Name

VkPipelineExecutableStatisticValueKHR - A union describing a pipeline
executable statistic



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineExecutableStatisticValueKHR` union is defined as:

``` c
// Provided by VK_KHR_pipeline_executable_properties
typedef union VkPipelineExecutableStatisticValueKHR {
    VkBool32    b32;
    int64_t     i64;
    uint64_t    u64;
    double      f64;
} VkPipelineExecutableStatisticValueKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `b32` is the 32-bit boolean value if the
  `VkPipelineExecutableStatisticFormatKHR` is
  `VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_BOOL32_KHR`.

- `i64` is the signed 64-bit integer value if the
  `VkPipelineExecutableStatisticFormatKHR` is
  `VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_INT64_KHR`.

- `u64` is the unsigned 64-bit integer value if the
  `VkPipelineExecutableStatisticFormatKHR` is
  `VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_UINT64_KHR`.

- `f64` is the 64-bit floating-point value if the
  `VkPipelineExecutableStatisticFormatKHR` is
  `VK_PIPELINE_EXECUTABLE_STATISTIC_FORMAT_FLOAT64_KHR`.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_pipeline_executable_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_executable_properties.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkPipelineExecutableStatisticKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableStatisticKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineExecutableStatisticValueKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
