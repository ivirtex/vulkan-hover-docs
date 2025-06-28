# VkImportMetalSharedEventInfoEXT(3) Manual Page

## Name

VkImportMetalSharedEventInfoEXT - Structure that identifies a VkSemaphore or VkEvent object and corresponding Metal Shared Event object to use.



## [](#_c_specification)C Specification

To import a Metal `id<MTLSharedEvent>` object to underlie a [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) or [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) object, include a `VkImportMetalSharedEventInfoEXT` structure in the `pNext` chain of the [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html) or [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateInfo.html) structure in a [vkCreateSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSemaphore.html) or [vkCreateEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateEvent.html) command, respectively.

The `VkImportMetalSharedEventInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_metal_objects
typedef struct VkImportMetalSharedEventInfoEXT {
    VkStructureType      sType;
    const void*          pNext;
    MTLSharedEvent_id    mtlSharedEvent;
} VkImportMetalSharedEventInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `mtlSharedEvent` is the Metal `id<MTLSharedEvent>` object that is to underlie the [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) or [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html).

## [](#_description)Description

If the `pNext` chain of the [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html) structure includes both `VkImportMetalSharedEventInfoEXT` and [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreTypeCreateInfo.html), the `signaledValue` property of the imported `id<MTLSharedEvent>` object will be set to [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreTypeCreateInfo.html)::`initialValue`.

Valid Usage (Implicit)

- [](#VUID-VkImportMetalSharedEventInfoEXT-sType-sType)VUID-VkImportMetalSharedEventInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_METAL_SHARED_EVENT_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportMetalSharedEventInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0