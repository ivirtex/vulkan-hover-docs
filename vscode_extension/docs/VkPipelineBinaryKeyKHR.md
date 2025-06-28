# VkPipelineBinaryKeyKHR(3) Manual Page

## Name

VkPipelineBinaryKeyKHR - Structure specifying a key to a pipeline binary



## [](#_c_specification)C Specification

The `VkPipelineBinaryKeyKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_binary
typedef struct VkPipelineBinaryKeyKHR {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           keySize;
    uint8_t            key[VK_MAX_PIPELINE_BINARY_KEY_SIZE_KHR];
} VkPipelineBinaryKeyKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `keySize` is the size, in bytes, of valid data returned in `key`.
- `key` is a buffer of opaque data specifying a pipeline binary key.

## [](#_description)Description

Any returned values beyond the first `keySize` bytes are undefined. Implementations **must** return a `keySize` greater than 0, and less-or-equal to `VK_MAX_PIPELINE_BINARY_KEY_SIZE_KHR`.

Two keys are considered equal if `keySize` is equal and the first `keySize` bytes of `key` compare equal.

Implementations **may** return a different `keySize` for different binaries.

Implementations **should** ensure that `keySize` is large enough to uniquely identify a pipeline binary.

Valid Usage (Implicit)

- [](#VUID-VkPipelineBinaryKeyKHR-sType-sType)VUID-VkPipelineBinaryKeyKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_BINARY_KEY_KHR`
- [](#VUID-VkPipelineBinaryKeyKHR-pNext-pNext)VUID-VkPipelineBinaryKeyKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkPipelineBinaryKeysAndDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeysAndDataKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPipelineBinaryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineBinaryDataKHR.html), [vkGetPipelineKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineKeyKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineBinaryKeyKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0