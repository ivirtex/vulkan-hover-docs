# VkPipelineLibraryCreateInfoKHR(3) Manual Page

## Name

VkPipelineLibraryCreateInfoKHR - Structure specifying pipeline libraries to use when creating a pipeline



## [](#_c_specification)C Specification

The `VkPipelineLibraryCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_library
typedef struct VkPipelineLibraryCreateInfoKHR {
    VkStructureType      sType;
    const void*          pNext;
    uint32_t             libraryCount;
    const VkPipeline*    pLibraries;
} VkPipelineLibraryCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `libraryCount` is the number of pipeline libraries in `pLibraries`.
- `pLibraries` is a pointer to an array of [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) structures specifying pipeline libraries to use when creating a pipeline.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-03381)VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-03381  
  Each element of `pLibraries` **must** have been created with `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`
- [](#VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-06855)VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-06855  
  If any library in `pLibraries` was created with a shader stage with [VkPipelineShaderStageModuleIdentifierCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageModuleIdentifierCreateInfoEXT.html) and `identifierSize` not equal to 0, the pipeline **must** be created with the `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` flag set
- [](#VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-08096)VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-08096  
  If any element of `pLibraries` was created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, all elements **must** have been created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07404)VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07404  
  If `pipeline` is being created with `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT`, every element of `pLibraries` **must** have been created with `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT`
- [](#VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07405)VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07405  
  If `pipeline` is being created without `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT`, every element of `pLibraries` **must** have been created without `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT`
- [](#VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07406)VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07406  
  If `pipeline` is being created with `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT`, every element of `pLibraries` **must** have been created with `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT`
- [](#VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07407)VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07407  
  If `pipeline` is being created without `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT`, every element of `pLibraries` **must** have been created without `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT`

Valid Usage (Implicit)

- [](#VUID-VkPipelineLibraryCreateInfoKHR-sType-sType)VUID-VkPipelineLibraryCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_LIBRARY_CREATE_INFO_KHR`
- [](#VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-parameter)VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-parameter  
  If `libraryCount` is not `0`, `pLibraries` **must** be a valid pointer to an array of `libraryCount` valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handles

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_library](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_library.html), [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineLibraryCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0