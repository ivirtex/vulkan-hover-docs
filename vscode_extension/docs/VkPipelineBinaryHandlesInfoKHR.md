# VkPipelineBinaryHandlesInfoKHR(3) Manual Page

## Name

VkPipelineBinaryHandlesInfoKHR - Structure containing newly created pipeline binaries



## [](#_c_specification)C Specification

The `VkPipelineBinaryHandlesInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_binary
typedef struct VkPipelineBinaryHandlesInfoKHR {
    VkStructureType         sType;
    const void*             pNext;
    uint32_t                pipelineBinaryCount;
    VkPipelineBinaryKHR*    pPipelineBinaries;
} VkPipelineBinaryHandlesInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pipelineBinaryCount` is the number of binaries associated with this pipeline or the number of entries in the `pPipelineBinaries` array.
- `pPipelineBinaries` is `NULL` or a pointer to an array of [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html) handles in which the resulting pipeline binaries are returned.

## [](#_description)Description

If `pPipelineBinaries` is `NULL`, the number of binaries that would be created is returned in `pipelineBinaryCount`. Otherwise, `pipelineBinaryCount` **must** be the number of entries in the `pPipelineBinaries` array, and on return from [vkCreatePipelineBinariesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePipelineBinariesKHR.html) `pipelineBinaryCount` is overwritten with the number of handles actually written to `pPipelineBinaries`. If the value of `pipelineBinaryCount` is less than the number of binaries that would have been created, at most `pipelineBinaryCount` handles will be written to `pPipelineBinaries` and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that `pPipelineBinaries` was not large enough to create all the binaries.

Valid Usage (Implicit)

- [](#VUID-VkPipelineBinaryHandlesInfoKHR-sType-sType)VUID-VkPipelineBinaryHandlesInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_BINARY_HANDLES_INFO_KHR`
- [](#VUID-VkPipelineBinaryHandlesInfoKHR-pNext-pNext)VUID-VkPipelineBinaryHandlesInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPipelineBinaryHandlesInfoKHR-pPipelineBinaries-parameter)VUID-VkPipelineBinaryHandlesInfoKHR-pPipelineBinaries-parameter  
  If `pipelineBinaryCount` is not `0`, and `pPipelineBinaries` is not `NULL`, `pPipelineBinaries` **must** be a valid pointer to an array of `pipelineBinaryCount` [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html) handles

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreatePipelineBinariesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePipelineBinariesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineBinaryHandlesInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0