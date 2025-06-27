# VkSparseImageMemoryBindInfo(3) Manual Page

## Name

VkSparseImageMemoryBindInfo - Structure specifying sparse image memory
bind information



## <a href="#_c_specification" class="anchor"></a>C Specification

Memory **can** be bound to sparse image blocks of `VkImage` objects
created with the `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` flag using the
following structure:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkSparseImageMemoryBindInfo {
    VkImage                           image;
    uint32_t                          bindCount;
    const VkSparseImageMemoryBind*    pBinds;
} VkSparseImageMemoryBindInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `image` is the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) object to be bound

- `bindCount` is the number of
  [VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryBind.html) structures in
  `pBinds` array

- `pBinds` is a pointer to an array of
  [VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryBind.html) structures

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSparseImageMemoryBindInfo-subresource-01722"
  id="VUID-VkSparseImageMemoryBindInfo-subresource-01722"></a>
  VUID-VkSparseImageMemoryBindInfo-subresource-01722  
  The `subresource.mipLevel` member of each element of `pBinds` **must**
  be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-VkSparseImageMemoryBindInfo-subresource-01723"
  id="VUID-VkSparseImageMemoryBindInfo-subresource-01723"></a>
  VUID-VkSparseImageMemoryBindInfo-subresource-01723  
  The `subresource.arrayLayer` member of each element of `pBinds`
  **must** be less than the `arrayLayers` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-VkSparseImageMemoryBindInfo-subresource-01106"
  id="VUID-VkSparseImageMemoryBindInfo-subresource-01106"></a>
  VUID-VkSparseImageMemoryBindInfo-subresource-01106  
  The `subresource.aspectMask` member of each element of `pBinds`
  **must** be valid for the `format` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-VkSparseImageMemoryBindInfo-image-02901"
  id="VUID-VkSparseImageMemoryBindInfo-image-02901"></a>
  VUID-VkSparseImageMemoryBindInfo-image-02901  
  `image` **must** have been created with
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` set

Valid Usage (Implicit)

- <a href="#VUID-VkSparseImageMemoryBindInfo-image-parameter"
  id="VUID-VkSparseImageMemoryBindInfo-image-parameter"></a>
  VUID-VkSparseImageMemoryBindInfo-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkSparseImageMemoryBindInfo-pBinds-parameter"
  id="VUID-VkSparseImageMemoryBindInfo-pBinds-parameter"></a>
  VUID-VkSparseImageMemoryBindInfo-pBinds-parameter  
  `pBinds` **must** be a valid pointer to an array of `bindCount` valid
  [VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryBind.html) structures

- <a href="#VUID-VkSparseImageMemoryBindInfo-bindCount-arraylength"
  id="VUID-VkSparseImageMemoryBindInfo-bindCount-arraylength"></a>
  VUID-VkSparseImageMemoryBindInfo-bindCount-arraylength  
  `bindCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindSparseInfo.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryBind.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSparseImageMemoryBindInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
