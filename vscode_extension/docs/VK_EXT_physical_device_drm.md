# VK\_EXT\_physical\_device\_drm(3) Manual Page

## Name

VK\_EXT\_physical\_device\_drm - device extension



## [](#_registered_extension_number)Registered Extension Number

354

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Simon Ser [\[GitHub\]emersion](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_physical_device_drm%5D%20%40emersion%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_physical_device_drm%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-06-09

**IP Status**

No known IP claims.

**Contributors**

- Simon Ser

## [](#_description)Description

This extension provides new facilities to query DRM properties for physical devices, enabling users to match Vulkan physical devices with DRM nodes on Linux.

Its functionality closely overlaps with `EGL_EXT_device_drm`[1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_physical_device_drm-fn1)^. Unlike the EGL extension, this extension does not expose a string containing the name of the device file and instead exposes device minor numbers.

DRM defines multiple device node types. Each physical device may have one primary node and one render node associated. Physical devices may have no primary node (e.g. if the device does not have a display subsystem), may have no render node (e.g. if it is a software rendering engine), or may have neither (e.g. if it is a software rendering engine without a display subsystem).

To query DRM properties for a physical device, chain [VkPhysicalDeviceDrmPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDrmPropertiesEXT.html) to [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html).

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceDrmPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDrmPropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_PHYSICAL_DEVICE_DRM_EXTENSION_NAME`
- `VK_EXT_PHYSICAL_DEVICE_DRM_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRM_PROPERTIES_EXT`

## [](#_references)References

1. []()[`EGL_EXT_device_drm`](https://registry.khronos.org/EGL/extensions/EXT/EGL_EXT_device_drm.txt)

## [](#_version_history)Version History

- Revision 1, 2021-06-09
  
  - First stable revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_physical_device_drm)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0