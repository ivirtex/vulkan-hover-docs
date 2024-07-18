# VkSparseBufferMemoryBindInfo(3) Manual Page

## Name

VkSparseBufferMemoryBindInfo - Structure specifying a sparse buffer
memory bind operation



## <a href="#_c_specification" class="anchor"></a>C Specification

Memory is bound to `VkBuffer` objects created with the
`VK_BUFFER_CREATE_SPARSE_BINDING_BIT` flag using the following
structure:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkSparseBufferMemoryBindInfo {
    VkBuffer                     buffer;
    uint32_t                     bindCount;
    const VkSparseMemoryBind*    pBinds;
} VkSparseBufferMemoryBindInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `buffer` is the [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) object to be bound.

- `bindCount` is the number of
  [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBind.html) structures in the
  `pBinds` array.

- `pBinds` is a pointer to an array of
  [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBind.html) structures.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkSparseBufferMemoryBindInfo-buffer-parameter"
  id="VUID-VkSparseBufferMemoryBindInfo-buffer-parameter"></a>
  VUID-VkSparseBufferMemoryBindInfo-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkSparseBufferMemoryBindInfo-pBinds-parameter"
  id="VUID-VkSparseBufferMemoryBindInfo-pBinds-parameter"></a>
  VUID-VkSparseBufferMemoryBindInfo-pBinds-parameter  
  `pBinds` **must** be a valid pointer to an array of `bindCount` valid
  [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBind.html) structures

- <a href="#VUID-VkSparseBufferMemoryBindInfo-bindCount-arraylength"
  id="VUID-VkSparseBufferMemoryBindInfo-bindCount-arraylength"></a>
  VUID-VkSparseBufferMemoryBindInfo-bindCount-arraylength  
  `bindCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindSparseInfo.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBind.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSparseBufferMemoryBindInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
