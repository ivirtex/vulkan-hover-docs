# VkPhysicalDeviceExternalFormatResolvePropertiesANDROID(3) Manual Page

## Name

VkPhysicalDeviceExternalFormatResolvePropertiesANDROID - Structure describing external format resolve supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceExternalFormatResolvePropertiesANDROID` structure is defined as:

```c++
// Provided by VK_ANDROID_external_format_resolve
typedef struct VkPhysicalDeviceExternalFormatResolvePropertiesANDROID {
    VkStructureType     sType;
    void*               pNext;
    VkBool32            nullColorAttachmentWithExternalFormatResolve;
    VkChromaLocation    externalFormatResolveChromaOffsetX;
    VkChromaLocation    externalFormatResolveChromaOffsetY;
} VkPhysicalDeviceExternalFormatResolvePropertiesANDROID;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`nullColorAttachmentWithExternalFormatResolve` indicates that there **must** be no color attachment image when performing external format resolves if it is `VK_TRUE`.
- []()`externalFormatResolveChromaOffsetX` indicates the [VkChromaLocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocation.html) that an implementation uses in the X axis for accesses to an external format image as a resolve attachment. This **must** be consistent between external format resolves and load operations from external format resolve attachments to color attachments when `nullColorAttachmentWithExternalFormatResolve` is `VK_TRUE`.
- []()`externalFormatResolveChromaOffsetY` indicates the [VkChromaLocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocation.html) that an implementation uses in the Y axis for accesses to an external format image as a resolve attachment. This **must** be consistent between external format resolves and load operations from external format resolve attachments to color attachments when `nullColorAttachmentWithExternalFormatResolve` is `VK_TRUE`.

## [](#_description)Description

If the `VkPhysicalDeviceExternalFormatResolvePropertiesANDROID` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceExternalFormatResolvePropertiesANDROID-sType-sType)VUID-VkPhysicalDeviceExternalFormatResolvePropertiesANDROID-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FORMAT_RESOLVE_PROPERTIES_ANDROID`

## [](#_see_also)See Also

[VK\_ANDROID\_external\_format\_resolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ANDROID_external_format_resolve.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkChromaLocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocation.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceExternalFormatResolvePropertiesANDROID).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0