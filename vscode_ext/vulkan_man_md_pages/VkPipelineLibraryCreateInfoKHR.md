# VkPipelineLibraryCreateInfoKHR(3) Manual Page

## Name

VkPipelineLibraryCreateInfoKHR - Structure specifying pipeline libraries
to use when creating a pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineLibraryCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_pipeline_library
typedef struct VkPipelineLibraryCreateInfoKHR {
    VkStructureType      sType;
    const void*          pNext;
    uint32_t             libraryCount;
    const VkPipeline*    pLibraries;
} VkPipelineLibraryCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `libraryCount` is the number of pipeline libraries in `pLibraries`.

- `pLibraries` is a pointer to an array of [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)
  structures specifying pipeline libraries to use when creating a
  pipeline.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-03381"
  id="VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-03381"></a>
  VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-03381  
  Each element of `pLibraries` **must** have been created with
  `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`

- <a href="#VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-06855"
  id="VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-06855"></a>
  VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-06855  
  If any library in `pLibraries` was created with a shader stage with
  [VkPipelineShaderStageModuleIdentifierCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageModuleIdentifierCreateInfoEXT.html)
  and `identifierSize` not equal to 0, the pipeline **must** be created
  with the `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT`
  flag set

- <a href="#VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-08096"
  id="VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-08096"></a>
  VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-08096  
  If any element of `pLibraries` was created with
  `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, all elements **must**
  have been created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07404"
  id="VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07404"></a>
  VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07404  
  If `pipeline` is being created with
  `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT_EXT`, every element of
  `pLibraries` **must** have been created with
  `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT_EXT`

- <a href="#VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07405"
  id="VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07405"></a>
  VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07405  
  If `pipeline` is being created without
  `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT_EXT`, every element of
  `pLibraries` **must** have been created without
  `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT_EXT`

- <a href="#VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07406"
  id="VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07406"></a>
  VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07406  
  If `pipeline` is being created with
  `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT_EXT`, every element of
  `pLibraries` **must** have been created with
  `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT_EXT`

- <a href="#VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07407"
  id="VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07407"></a>
  VUID-VkPipelineLibraryCreateInfoKHR-pipeline-07407  
  If `pipeline` is being created without
  `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT_EXT`, every element of
  `pLibraries` **must** have been created without
  `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineLibraryCreateInfoKHR-sType-sType"
  id="VUID-VkPipelineLibraryCreateInfoKHR-sType-sType"></a>
  VUID-VkPipelineLibraryCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_LIBRARY_CREATE_INFO_KHR`

- <a href="#VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-parameter"
  id="VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-parameter"></a>
  VUID-VkPipelineLibraryCreateInfoKHR-pLibraries-parameter  
  If `libraryCount` is not `0`, `pLibraries` **must** be a valid pointer
  to an array of `libraryCount` valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)
  handles

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_pipeline_library](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_library.html),
[VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineLibraryCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
