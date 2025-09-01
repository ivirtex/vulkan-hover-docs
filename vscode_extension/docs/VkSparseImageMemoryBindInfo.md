# VkSparseImageMemoryBindInfo(3) Manual Page

## Name

VkSparseImageMemoryBindInfo - Structure specifying sparse image memory bind information



## [](#_c_specification)C Specification

Memory **can** be bound to sparse image blocks of `VkImage` objects created with the `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` flag using the following structure:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkSparseImageMemoryBindInfo {
    VkImage                           image;
    uint32_t                          bindCount;
    const VkSparseImageMemoryBind*    pBinds;
} VkSparseImageMemoryBindInfo;
```

## [](#_members)Members

- `image` is the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) object to be bound
- `bindCount` is the number of [VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryBind.html) structures in `pBinds` array
- `pBinds` is a pointer to an array of [VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryBind.html) structures

## [](#_description)Description

Valid Usage

- [](#VUID-VkSparseImageMemoryBindInfo-subresource-01722)VUID-VkSparseImageMemoryBindInfo-subresource-01722  
  The `subresource.mipLevel` member of each element of `pBinds` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkSparseImageMemoryBindInfo-subresource-01723)VUID-VkSparseImageMemoryBindInfo-subresource-01723  
  The `subresource.arrayLayer` member of each element of `pBinds` **must** be less than the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkSparseImageMemoryBindInfo-subresource-01106)VUID-VkSparseImageMemoryBindInfo-subresource-01106  
  The `subresource.aspectMask` member of each element of `pBinds` **must** be valid for the `format` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkSparseImageMemoryBindInfo-image-02901)VUID-VkSparseImageMemoryBindInfo-image-02901  
  `image` **must** have been created with `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` set

Valid Usage (Implicit)

- [](#VUID-VkSparseImageMemoryBindInfo-image-parameter)VUID-VkSparseImageMemoryBindInfo-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkSparseImageMemoryBindInfo-pBinds-parameter)VUID-VkSparseImageMemoryBindInfo-pBinds-parameter  
  `pBinds` **must** be a valid pointer to an array of `bindCount` valid [VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryBind.html) structures
- [](#VUID-VkSparseImageMemoryBindInfo-bindCount-arraylength)VUID-VkSparseImageMemoryBindInfo-bindCount-arraylength  
  `bindCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryBind.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSparseImageMemoryBindInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0