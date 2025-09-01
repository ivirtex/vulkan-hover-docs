# VkPhysicalDeviceDrmPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceDrmPropertiesEXT - Structure containing DRM information of a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDrmPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_physical_device_drm
typedef struct VkPhysicalDeviceDrmPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           hasPrimary;
    VkBool32           hasRender;
    int64_t            primaryMajor;
    int64_t            primaryMinor;
    int64_t            renderMajor;
    int64_t            renderMinor;
} VkPhysicalDeviceDrmPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `hasPrimary` is a boolean indicating whether the physical device has a DRM primary node.
- `hasRender` is a boolean indicating whether the physical device has a DRM render node.
- `primaryMajor` is the DRM primary node major number, if any.
- `primaryMinor` is the DRM primary node minor number, if any.
- `renderMajor` is the DRM render node major number, if any.
- `renderMinor` is the DRM render node minor number, if any.

## [](#_description)Description

If the `VkPhysicalDeviceDrmPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These are properties of the DRM information of a physical device.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDrmPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceDrmPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRM_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_physical\_device\_drm](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_physical_device_drm.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDrmPropertiesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0