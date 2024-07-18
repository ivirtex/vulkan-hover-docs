# VkExportMetalSharedEventInfoEXT(3) Manual Page

## Name

VkExportMetalSharedEventInfoEXT - Structure that identifies a
VkSemaphore or VkEvent object and corresponding Metal MTLSharedEvent
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To export the Metal `MTLSharedEvent` object underlying a
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) or [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) object,
include a `VkExportMetalSharedEventInfoEXT` structure in the `pNext`
chain of the `pMetalObjectsInfo` parameter of a
[vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkExportMetalObjectsEXT.html) call.

The `VkExportMetalSharedEventInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalSharedEventInfoEXT {
    VkStructureType      sType;
    const void*          pNext;
    VkSemaphore          semaphore;
    VkEvent              event;
    MTLSharedEvent_id    mtlSharedEvent;
} VkExportMetalSharedEventInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `semaphore` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a
  [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html).

- `event` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a
  [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html).

- `mtlSharedEvent` is the Metal `id<MTLSharedEvent>` object underlying
  the [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) or [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) object
  in `semaphore` or `event`, respectively. The implementation will
  return the `MTLSharedEvent` in this member, or it will return `NULL`
  if no `MTLSharedEvent` could be found underlying the
  [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) or [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkExportMetalSharedEventInfoEXT-sType-sType"
  id="VUID-VkExportMetalSharedEventInfoEXT-sType-sType"></a>
  VUID-VkExportMetalSharedEventInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_EXPORT_METAL_SHARED_EVENT_INFO_EXT`

- <a href="#VUID-VkExportMetalSharedEventInfoEXT-semaphore-parameter"
  id="VUID-VkExportMetalSharedEventInfoEXT-semaphore-parameter"></a>
  VUID-VkExportMetalSharedEventInfoEXT-semaphore-parameter  
  If `semaphore` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

- <a href="#VUID-VkExportMetalSharedEventInfoEXT-event-parameter"
  id="VUID-VkExportMetalSharedEventInfoEXT-event-parameter"></a>
  VUID-VkExportMetalSharedEventInfoEXT-event-parameter  
  If `event` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `event`
  **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) handle

- <a href="#VUID-VkExportMetalSharedEventInfoEXT-commonparent"
  id="VUID-VkExportMetalSharedEventInfoEXT-commonparent"></a>
  VUID-VkExportMetalSharedEventInfoEXT-commonparent  
  Both of `event`, and `semaphore` that are valid handles of non-ignored
  parameters **must** have been created, allocated, or retrieved from
  the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMetalSharedEventInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
