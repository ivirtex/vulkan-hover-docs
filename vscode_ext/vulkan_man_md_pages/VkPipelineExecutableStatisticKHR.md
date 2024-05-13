# VkPipelineExecutableStatisticKHR(3) Manual Page

## Name

VkPipelineExecutableStatisticKHR - Structure describing a compile time
pipeline executable statistic



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineExecutableStatisticKHR` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `name` is an array of `VK_MAX_DESCRIPTION_SIZE` `char` containing a
  null-terminated UTF-8 string which is a short human readable name for
  this statistic.

- `description` is an array of `VK_MAX_DESCRIPTION_SIZE` `char`
  containing a null-terminated UTF-8 string which is a human readable
  description for this statistic.

- `format` is a
  [VkPipelineExecutableStatisticFormatKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableStatisticFormatKHR.html)
  value specifying the format of the data found in `value`.

- `value` is the value of this statistic.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineExecutableStatisticKHR-sType-sType"
  id="VUID-VkPipelineExecutableStatisticKHR-sType-sType"></a>
  VUID-VkPipelineExecutableStatisticKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_STATISTIC_KHR`

- <a href="#VUID-VkPipelineExecutableStatisticKHR-pNext-pNext"
  id="VUID-VkPipelineExecutableStatisticKHR-pNext-pNext"></a>
  VUID-VkPipelineExecutableStatisticKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_pipeline_executable_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_executable_properties.html),
[VkPipelineExecutableStatisticFormatKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableStatisticFormatKHR.html),
[VkPipelineExecutableStatisticValueKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableStatisticValueKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPipelineExecutableStatisticsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutableStatisticsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineExecutableStatisticKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
