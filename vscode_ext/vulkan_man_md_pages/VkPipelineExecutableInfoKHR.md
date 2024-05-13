# VkPipelineExecutableInfoKHR(3) Manual Page

## Name

VkPipelineExecutableInfoKHR - Structure describing a pipeline executable
to query for associated statistics or internal representations



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineExecutableInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_pipeline_executable_properties
typedef struct VkPipelineExecutableInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkPipeline         pipeline;
    uint32_t           executableIndex;
} VkPipelineExecutableInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pipeline` is the pipeline to query.

- `executableIndex` is the index of the pipeline executable to query in
  the array of executable properties returned by
  [vkGetPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutablePropertiesKHR.html).

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPipelineExecutableInfoKHR-executableIndex-03275"
  id="VUID-VkPipelineExecutableInfoKHR-executableIndex-03275"></a>
  VUID-VkPipelineExecutableInfoKHR-executableIndex-03275  
  `executableIndex` **must** be less than the number of pipeline
  executables associated with `pipeline` as returned in the
  `pExecutableCount` parameter of `vkGetPipelineExecutablePropertiesKHR`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineExecutableInfoKHR-sType-sType"
  id="VUID-VkPipelineExecutableInfoKHR-sType-sType"></a>
  VUID-VkPipelineExecutableInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INFO_KHR`

- <a href="#VUID-VkPipelineExecutableInfoKHR-pNext-pNext"
  id="VUID-VkPipelineExecutableInfoKHR-pNext-pNext"></a>
  VUID-VkPipelineExecutableInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkPipelineExecutableInfoKHR-pipeline-parameter"
  id="VUID-VkPipelineExecutableInfoKHR-pipeline-parameter"></a>
  VUID-VkPipelineExecutableInfoKHR-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_pipeline_executable_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_executable_properties.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPipelineExecutableInternalRepresentationsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutableInternalRepresentationsKHR.html),
[vkGetPipelineExecutableStatisticsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutableStatisticsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineExecutableInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
