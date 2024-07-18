# VkPipelineInfoKHR(3) Manual Page

## Name

VkPipelineInfoKHR - Structure describing a pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_pipeline_executable_properties
typedef struct VkPipelineInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkPipeline         pipeline;
} VkPipelineInfoKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_pipeline_properties
typedef VkPipelineInfoKHR VkPipelineInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pipeline` is a `VkPipeline` handle.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineInfoKHR-sType-sType"
  id="VUID-VkPipelineInfoKHR-sType-sType"></a>
  VUID-VkPipelineInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_INFO_KHR`

- <a href="#VUID-VkPipelineInfoKHR-pNext-pNext"
  id="VUID-VkPipelineInfoKHR-pNext-pNext"></a>
  VUID-VkPipelineInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkPipelineInfoKHR-pipeline-parameter"
  id="VUID-VkPipelineInfoKHR-pipeline-parameter"></a>
  VUID-VkPipelineInfoKHR-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pipeline_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_properties.html),
[VK_KHR_pipeline_executable_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_executable_properties.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutablePropertiesKHR.html),
[vkGetPipelinePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelinePropertiesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
