# VkImportMetalSharedEventInfoEXT(3) Manual Page

## Name

VkImportMetalSharedEventInfoEXT - Structure that identifies a
VkSemaphore or VkEvent object and corresponding Metal Shared Event
object to use.



## <a href="#_c_specification" class="anchor"></a>C Specification

To import a Metal `id<MTLSharedEvent>` object to underlie a
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) or [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) object,
include a `VkImportMetalSharedEventInfoEXT` structure in the `pNext`
chain of the [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html) or
[VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateInfo.html) structure in a
[vkCreateSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSemaphore.html) or
[vkCreateEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateEvent.html) command, respectively.

The `VkImportMetalSharedEventInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_metal_objects
typedef struct VkImportMetalSharedEventInfoEXT {
    VkStructureType      sType;
    const void*          pNext;
    MTLSharedEvent_id    mtlSharedEvent;
} VkImportMetalSharedEventInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `mtlSharedEvent` is the Metal `id<MTLSharedEvent>` object that is to
  underlie the [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) or
  [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html).

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of the
[VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html) structure includes
both `VkImportMetalSharedEventInfoEXT` and
[VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html), the
`signaledValue` property of the imported `id<MTLSharedEvent>` object
will be set to `initialValue` of
[VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html).

Valid Usage (Implicit)

- <a href="#VUID-VkImportMetalSharedEventInfoEXT-sType-sType"
  id="VUID-VkImportMetalSharedEventInfoEXT-sType-sType"></a>
  VUID-VkImportMetalSharedEventInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMPORT_METAL_SHARED_EVENT_INFO_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportMetalSharedEventInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
