# VkPhysicalDeviceDrmPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceDrmPropertiesEXT - Structure containing DRM information
of a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDrmPropertiesEXT` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `hasPrimary` is a boolean indicating whether the physical device has a
  DRM primary node.

- `hasRender` is a boolean indicating whether the physical device has a
  DRM render node.

- `primaryMajor` is the DRM primary node major number, if any.

- `primaryMinor` is the DRM primary node minor number, if any.

- `renderMajor` is the DRM render node major number, if any.

- `renderMinor` is the DRM render node minor number, if any.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceDrmPropertiesEXT` structure is included in the
`pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

These are properties of the DRM information of a physical device.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceDrmPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceDrmPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceDrmPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRM_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_physical_device_drm](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_physical_device_drm.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDrmPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
