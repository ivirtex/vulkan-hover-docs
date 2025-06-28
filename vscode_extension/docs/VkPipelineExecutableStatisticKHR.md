# VkPipelineExecutableStatisticKHR(3) Manual Page

## Name

VkPipelineExecutableStatisticKHR - Structure describing a compile time pipeline executable statistic



## [](#_c_specification)C Specification

The `VkPipelineExecutableStatisticKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_executable_properties
typedef struct VkPipelineExecutableStatisticKHR {
    VkStructureType                           sType;
    void*                                     pNext;
    char                                      name[VK_MAX_DESCRIPTION_SIZE];
    char                                      description[VK_MAX_DESCRIPTION_SIZE];
    VkPipelineExecutableStatisticFormatKHR    format;
    VkPipelineExecutableStatisticValueKHR     value;
} VkPipelineExecutableStatisticKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `name` is an array of `VK_MAX_DESCRIPTION_SIZE` `char` containing a null-terminated UTF-8 string which is a short human readable name for this statistic.
- `description` is an array of `VK_MAX_DESCRIPTION_SIZE` `char` containing a null-terminated UTF-8 string which is a human readable description for this statistic.
- `format` is a [VkPipelineExecutableStatisticFormatKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutableStatisticFormatKHR.html) value specifying the format of the data found in `value`.
- `value` is the value of this statistic.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPipelineExecutableStatisticKHR-sType-sType)VUID-VkPipelineExecutableStatisticKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_STATISTIC_KHR`
- [](#VUID-VkPipelineExecutableStatisticKHR-pNext-pNext)VUID-VkPipelineExecutableStatisticKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_executable\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_executable_properties.html), [VkPipelineExecutableStatisticFormatKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutableStatisticFormatKHR.html), [VkPipelineExecutableStatisticValueKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutableStatisticValueKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPipelineExecutableStatisticsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineExecutableStatisticsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineExecutableStatisticKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0