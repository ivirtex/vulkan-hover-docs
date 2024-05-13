# VkSparseImageOpaqueMemoryBindInfo(3) Manual Page

## Name

VkSparseImageOpaqueMemoryBindInfo - Structure specifying sparse image
opaque memory bind information



## <a href="#_c_specification" class="anchor"></a>C Specification

Memory is bound to opaque regions of `VkImage` objects created with the
`VK_IMAGE_CREATE_SPARSE_BINDING_BIT` flag using the following structure:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkSparseImageOpaqueMemoryBindInfo {
    VkImage                      image;
    uint32_t                     bindCount;
    const VkSparseMemoryBind*    pBinds;
} VkSparseImageOpaqueMemoryBindInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `image` is the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) object to be bound.

- `bindCount` is the number of
  [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBind.html) structures in the
  `pBinds` array.

- `pBinds` is a pointer to an array of
  [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBind.html) structures.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSparseImageOpaqueMemoryBindInfo-pBinds-01103"
  id="VUID-VkSparseImageOpaqueMemoryBindInfo-pBinds-01103"></a>
  VUID-VkSparseImageOpaqueMemoryBindInfo-pBinds-01103  
  If the `flags` member of any element of `pBinds` contains
  `VK_SPARSE_MEMORY_BIND_METADATA_BIT`, the binding range defined
  **must** be within the mip tail region of the metadata aspect of
  `image`

Valid Usage (Implicit)

- <a href="#VUID-VkSparseImageOpaqueMemoryBindInfo-image-parameter"
  id="VUID-VkSparseImageOpaqueMemoryBindInfo-image-parameter"></a>
  VUID-VkSparseImageOpaqueMemoryBindInfo-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkSparseImageOpaqueMemoryBindInfo-pBinds-parameter"
  id="VUID-VkSparseImageOpaqueMemoryBindInfo-pBinds-parameter"></a>
  VUID-VkSparseImageOpaqueMemoryBindInfo-pBinds-parameter  
  `pBinds` **must** be a valid pointer to an array of `bindCount` valid
  [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBind.html) structures

- <a href="#VUID-VkSparseImageOpaqueMemoryBindInfo-bindCount-arraylength"
  id="VUID-VkSparseImageOpaqueMemoryBindInfo-bindCount-arraylength"></a>
  VUID-VkSparseImageOpaqueMemoryBindInfo-bindCount-arraylength  
  `bindCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindSparseInfo.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBind.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSparseImageOpaqueMemoryBindInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
