# VkExternalMemoryAcquireUnmodifiedEXT(3) Manual Page

## Name

VkExternalMemoryAcquireUnmodifiedEXT - Structure specifying that external memory has remained unmodified since releasing ownership



## [](#_c_specification)C Specification

An *acquire operation* **may** have a performance penalty when acquiring ownership of a subresource range from one of the special queue families reserved for external memory ownership transfers described above. The application **can** reduce the performance penalty in some cases by adding a [VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html) structure to the `pNext` chain of the *acquire operation*'s memory barrier structure.

The `VkExternalMemoryAcquireUnmodifiedEXT` structure is defined as:

```c++
// Provided by VK_EXT_external_memory_acquire_unmodified
typedef struct VkExternalMemoryAcquireUnmodifiedEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           acquireUnmodifiedMemory;
} VkExternalMemoryAcquireUnmodifiedEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `acquireUnmodifiedMemory` specifies, if `VK_TRUE`, that no range of [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) bound to the resource of the memory barrier’s subresource range was modified at any time since the resource’s most recent release of ownership to the queue family specified by the memory barrier’s `srcQueueFamilyIndex`. If `VK_FALSE`, it specifies nothing.

## [](#_description)Description

If the application releases ownership of the subresource range to one of the special queue families reserved for external memory ownership transfers with a memory barrier structure, and later re-acquires ownership from the same queue family with a memory barrier structure, and if no range of [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) bound to the resource was modified at any time between the *release operation* and the *acquire operation*, then the application **should** add a [VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html) structure to the `pNext` chain of the *acquire operation*'s memory barrier structure because this **may** reduce the performance penalty.

This structure is ignored if `acquireUnmodifiedMemory` is `VK_FALSE`. In particular, `VK_FALSE` does *not* specify that memory was modified.

This structure is ignored if the memory barrier’s `srcQueueFamilyIndex` is not a special queue family reserved for external memory ownership transfers.

Note

The method by which the application determines whether memory was modified between the *release operation* and *acquire operation* is outside the scope of Vulkan.

For any Vulkan operation that accesses a resource, the application **must** not assume the implementation accesses the resource’s memory as read-only, even for *apparently* read-only operations such as transfer commands and shader reads.

The validity of [VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html)::`acquireUnmodifiedMemory` is independent of memory ranges outside the ranges of [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) bound to the resource. In particular, it is independent of any implementation-private memory associated with the resource.

Valid Usage

- [](#VUID-VkExternalMemoryAcquireUnmodifiedEXT-acquireUnmodifiedMemory-08922)VUID-VkExternalMemoryAcquireUnmodifiedEXT-acquireUnmodifiedMemory-08922  
  If `acquireUnmodifiedMemory` is `VK_TRUE`, and the memory barrier’s `srcQueueFamilyIndex` is a special queue family reserved for external memory ownership transfers (as described in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers)), then each range of [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) bound to the resource **must** have remained unmodified during all time since the resource’s most recent release of ownership to the queue family

Valid Usage (Implicit)

- [](#VUID-VkExternalMemoryAcquireUnmodifiedEXT-sType-sType)VUID-VkExternalMemoryAcquireUnmodifiedEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_ACQUIRE_UNMODIFIED_EXT`

## [](#_see_also)See Also

[VK\_EXT\_external\_memory\_acquire\_unmodified](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_external_memory_acquire_unmodified.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalMemoryAcquireUnmodifiedEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0