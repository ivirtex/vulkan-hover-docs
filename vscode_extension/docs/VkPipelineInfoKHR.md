# VkPipelineInfoKHR(3) Manual Page

## Name

VkPipelineInfoKHR - Structure describing a pipeline



## [](#_c_specification)C Specification

The `VkPipelineInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_executable_properties
typedef struct VkPipelineInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkPipeline         pipeline;
} VkPipelineInfoKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_pipeline_properties
typedef VkPipelineInfoKHR VkPipelineInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pipeline` is a `VkPipeline` handle.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPipelineInfoKHR-sType-sType)VUID-VkPipelineInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_INFO_KHR`
- [](#VUID-VkPipelineInfoKHR-pNext-pNext)VUID-VkPipelineInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPipelineInfoKHR-pipeline-parameter)VUID-VkPipelineInfoKHR-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle

## [](#_see_also)See Also

[VK\_EXT\_pipeline\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_properties.html), [VK\_KHR\_pipeline\_executable\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_executable_properties.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineExecutablePropertiesKHR.html), [vkGetPipelinePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelinePropertiesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0