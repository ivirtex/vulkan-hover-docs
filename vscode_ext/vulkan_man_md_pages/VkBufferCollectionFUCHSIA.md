# VkBufferCollectionFUCHSIA(3) Manual Page

## Name

VkBufferCollectionFUCHSIA - Opaque handle to a buffer collection object



## <a href="#_c_specification" class="anchor"></a>C Specification

Fuchsia’s FIDL-based Sysmem service interoperates with Vulkan via the
[`VK_FUCHSIA_buffer_collection`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html)
extension.

A buffer collection is a set of one or more buffers which were allocated
together as a group and which all have the same properties. These
properties describe the buffers' internal representation, such as its
dimensions and memory layout. This ensures that all of the buffers can
be used interchangeably by tasks that require swapping among multiple
buffers, such as double-buffered graphics rendering.

On Fuchsia, the Sysmem service uses buffer collections as a core
construct in its design.

Buffer collections are represented by `VkBufferCollectionFUCHSIA`
handles:

``` c
// Provided by VK_FUCHSIA_buffer_collection
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkBufferCollectionFUCHSIA)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html),
[VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html),
[VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryBufferCollectionFUCHSIA.html),
[vkCreateBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateBufferCollectionFUCHSIA.html),
[vkDestroyBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyBufferCollectionFUCHSIA.html),
[vkGetBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferCollectionPropertiesFUCHSIA.html),
[vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html),
[vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferCollectionFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
