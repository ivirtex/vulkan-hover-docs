# VkBufferCollectionFUCHSIA(3) Manual Page

## Name

VkBufferCollectionFUCHSIA - Opaque handle to a buffer collection object



## [](#_c_specification)C Specification

Fuchsia’s FIDL-based Sysmem service interoperates with Vulkan via the `VK_FUCHSIA_buffer_collection` extension.

A buffer collection is a set of one or more buffers which were allocated together as a group and which all have the same properties. These properties describe the buffers' internal representation, such as its dimensions and memory layout. This ensures that all of the buffers can be used interchangeably by tasks that require swapping among multiple buffers, such as double-buffered graphics rendering.

On Fuchsia, the Sysmem service uses buffer collections as a core construct in its design.

Buffer collections are represented by `VkBufferCollectionFUCHSIA` handles:

```c++
// Provided by VK_FUCHSIA_buffer_collection
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkBufferCollectionFUCHSIA)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html), [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html), [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryBufferCollectionFUCHSIA.html), [vkCreateBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateBufferCollectionFUCHSIA.html), [vkDestroyBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyBufferCollectionFUCHSIA.html), [vkGetBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferCollectionPropertiesFUCHSIA.html), [vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html), [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferCollectionFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0