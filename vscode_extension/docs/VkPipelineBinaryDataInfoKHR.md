# VkPipelineBinaryDataInfoKHR(3) Manual Page

## Name

VkPipelineBinaryDataInfoKHR - Structure specifying a pipeline binary to retrieve binary data from



## [](#_c_specification)C Specification

The `VkPipelineBinaryDataInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_binary
typedef struct VkPipelineBinaryDataInfoKHR {
    VkStructureType        sType;
    void*                  pNext;
    VkPipelineBinaryKHR    pipelineBinary;
} VkPipelineBinaryDataInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pipelineBinary` is the pipeline binary to get data from.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPipelineBinaryDataInfoKHR-sType-sType)VUID-VkPipelineBinaryDataInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_BINARY_DATA_INFO_KHR`
- [](#VUID-VkPipelineBinaryDataInfoKHR-pNext-pNext)VUID-VkPipelineBinaryDataInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPipelineBinaryDataInfoKHR-pipelineBinary-parameter)VUID-VkPipelineBinaryDataInfoKHR-pipelineBinary-parameter  
  `pipelineBinary` **must** be a valid [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html) handle

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPipelineBinaryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineBinaryDataKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineBinaryDataInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0