# VkSparseBufferMemoryBindInfo(3) Manual Page

## Name

VkSparseBufferMemoryBindInfo - Structure specifying a sparse buffer memory bind operation



## [](#_c_specification)C Specification

Memory is bound to `VkBuffer` objects created with the `VK_BUFFER_CREATE_SPARSE_BINDING_BIT` flag using the following structure:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkSparseBufferMemoryBindInfo {
    VkBuffer                     buffer;
    uint32_t                     bindCount;
    const VkSparseMemoryBind*    pBinds;
} VkSparseBufferMemoryBindInfo;
```

## [](#_members)Members

- `buffer` is the [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) object to be bound.
- `bindCount` is the number of [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBind.html) structures in the `pBinds` array.
- `pBinds` is a pointer to an array of [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBind.html) structures.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSparseBufferMemoryBindInfo-buffer-parameter)VUID-VkSparseBufferMemoryBindInfo-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkSparseBufferMemoryBindInfo-pBinds-parameter)VUID-VkSparseBufferMemoryBindInfo-pBinds-parameter  
  `pBinds` **must** be a valid pointer to an array of `bindCount` valid [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBind.html) structures
- [](#VUID-VkSparseBufferMemoryBindInfo-bindCount-arraylength)VUID-VkSparseBufferMemoryBindInfo-bindCount-arraylength  
  `bindCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBind.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSparseBufferMemoryBindInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0