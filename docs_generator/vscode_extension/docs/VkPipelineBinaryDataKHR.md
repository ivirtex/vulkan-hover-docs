# VkPipelineBinaryDataKHR(3) Manual Page

## Name

VkPipelineBinaryDataKHR - Structure specifying data and length of a pipeline binary



## [](#_c_specification)C Specification

The `VkPipelineBinaryDataKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_binary
typedef struct VkPipelineBinaryDataKHR {
    size_t    dataSize;
    void*     pData;
} VkPipelineBinaryDataKHR;
```

## [](#_members)Members

- `dataSize` is the size of the `pData` buffer in bytes.
- `pData` is a pointer to a buffer of `size` bytes that contains pipeline binary data obtained from `vkGetPipelineBinaryDataKHR`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPipelineBinaryDataKHR-pData-parameter)VUID-VkPipelineBinaryDataKHR-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes
- [](#VUID-VkPipelineBinaryDataKHR-dataSize-arraylength)VUID-VkPipelineBinaryDataKHR-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkPipelineBinaryKeysAndDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeysAndDataKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineBinaryDataKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0