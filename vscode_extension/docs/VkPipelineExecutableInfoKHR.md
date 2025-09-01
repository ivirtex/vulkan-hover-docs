# VkPipelineExecutableInfoKHR(3) Manual Page

## Name

VkPipelineExecutableInfoKHR - Structure describing a pipeline executable to query for associated statistics or internal representations



## [](#_c_specification)C Specification

The `VkPipelineExecutableInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_executable_properties
typedef struct VkPipelineExecutableInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkPipeline         pipeline;
    uint32_t           executableIndex;
} VkPipelineExecutableInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pipeline` is the pipeline to query.
- `executableIndex` is the index of the pipeline executable to query in the array of executable properties returned by [vkGetPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineExecutablePropertiesKHR.html).

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineExecutableInfoKHR-executableIndex-03275)VUID-VkPipelineExecutableInfoKHR-executableIndex-03275  
  `executableIndex` **must** be less than the number of pipeline executables associated with `pipeline` as returned in the `pExecutableCount` parameter of `vkGetPipelineExecutablePropertiesKHR`

Valid Usage (Implicit)

- [](#VUID-VkPipelineExecutableInfoKHR-sType-sType)VUID-VkPipelineExecutableInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INFO_KHR`
- [](#VUID-VkPipelineExecutableInfoKHR-pNext-pNext)VUID-VkPipelineExecutableInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPipelineExecutableInfoKHR-pipeline-parameter)VUID-VkPipelineExecutableInfoKHR-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_executable\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_executable_properties.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPipelineExecutableInternalRepresentationsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineExecutableInternalRepresentationsKHR.html), [vkGetPipelineExecutableStatisticsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineExecutableStatisticsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineExecutableInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0