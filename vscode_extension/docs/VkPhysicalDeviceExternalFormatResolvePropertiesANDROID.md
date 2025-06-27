# VkPhysicalDeviceExternalFormatResolvePropertiesANDROID(3) Manual Page

## Name

VkPhysicalDeviceExternalFormatResolvePropertiesANDROID - Structure
describing external format resolve supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceExternalFormatResolvePropertiesANDROID` structure
is defined as:

``` c
// Provided by VK_ANDROID_external_format_resolve
typedef struct VkPhysicalDeviceExternalFormatResolvePropertiesANDROID {
    VkStructureType     sType;
    void*               pNext;
    VkBool32            nullColorAttachmentWithExternalFormatResolve;
    VkChromaLocation    externalFormatResolveChromaOffsetX;
    VkChromaLocation    externalFormatResolveChromaOffsetY;
} VkPhysicalDeviceExternalFormatResolvePropertiesANDROID;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-nullColorAttachmentWithExternalFormatResolve"></span>
  `nullColorAttachmentWithExternalFormatResolve` indicates that there
  **must** be no color attachment image when performing external format
  resolves if it is `VK_TRUE`.

- <span id="limits-externalFormatResolveChromaOffsetX"></span>
  `externalFormatResolveChromaOffsetX` indicates the
  [VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html) that an implementation uses
  in the X axis for accesses to an external format image as a resolve
  attachment. This **must** be consistent between external format
  resolves and load operations from external format resolve attachments
  to color attachments when
  `nullColorAttachmentWithExternalFormatResolve` is `VK_TRUE`.

- <span id="limits-externalFormatResolveChromaOffsetY"></span>
  `externalFormatResolveChromaOffsetY` indicates the
  [VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html) that an implementation uses
  in the Y axis for accesses to an external format image as a resolve
  attachment. This **must** be consistent between external format
  resolves and load operations from external format resolve attachments
  to color attachments when
  `nullColorAttachmentWithExternalFormatResolve` is `VK_TRUE`.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceExternalFormatResolvePropertiesANDROID`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceExternalFormatResolvePropertiesANDROID-sType-sType"
  id="VUID-VkPhysicalDeviceExternalFormatResolvePropertiesANDROID-sType-sType"></a>
  VUID-VkPhysicalDeviceExternalFormatResolvePropertiesANDROID-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FORMAT_RESOLVE_PROPERTIES_ANDROID`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ANDROID_external_format_resolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_format_resolve.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceExternalFormatResolvePropertiesANDROID"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
