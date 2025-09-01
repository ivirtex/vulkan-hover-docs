# VkExportMetalSharedEventInfoEXT(3) Manual Page

## Name

VkExportMetalSharedEventInfoEXT - Structure that identifies a VkSemaphore or VkEvent object and corresponding Metal MTLSharedEvent object



## [](#_c_specification)C Specification

To export the Metal `MTLSharedEvent` object underlying a [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) or [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) object, include a `VkExportMetalSharedEventInfoEXT` structure in the `pNext` chain of the `pMetalObjectsInfo` parameter of a [vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkExportMetalObjectsEXT.html) call.

The `VkExportMetalSharedEventInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalSharedEventInfoEXT {
    VkStructureType      sType;
    const void*          pNext;
    VkSemaphore          semaphore;
    VkEvent              event;
    MTLSharedEvent_id    mtlSharedEvent;
} VkExportMetalSharedEventInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `semaphore` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html).
- `event` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html).
- `mtlSharedEvent` is the Metal `id<MTLSharedEvent>` object underlying the [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) or [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) object in `semaphore` or `event`, respectively. The implementation will return the `MTLSharedEvent` in this member, or it will return `NULL` if no `MTLSharedEvent` could be found underlying the [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) or [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) object.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExportMetalSharedEventInfoEXT-sType-sType)VUID-VkExportMetalSharedEventInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_METAL_SHARED_EVENT_INFO_EXT`
- [](#VUID-VkExportMetalSharedEventInfoEXT-semaphore-parameter)VUID-VkExportMetalSharedEventInfoEXT-semaphore-parameter  
  If `semaphore` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handle
- [](#VUID-VkExportMetalSharedEventInfoEXT-event-parameter)VUID-VkExportMetalSharedEventInfoEXT-event-parameter  
  If `event` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) handle
- [](#VUID-VkExportMetalSharedEventInfoEXT-commonparent)VUID-VkExportMetalSharedEventInfoEXT-commonparent  
  Both of `event`, and `semaphore` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportMetalSharedEventInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0