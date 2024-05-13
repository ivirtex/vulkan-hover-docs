# VK_EXT_physical_device_drm(3) Manual Page

## Name

VK_EXT_physical_device_drm - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

354

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Simon Ser <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_physical_device_drm%5D%20@emersion%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_physical_device_drm%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>emersion</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-06-09

**IP Status**  
No known IP claims.

**Contributors**  
- Simon Ser

## <a href="#_description" class="anchor"></a>Description

This extension provides new facilities to query DRM properties for
physical devices, enabling users to match Vulkan physical devices with
DRM nodes on Linux.

Its functionality closely overlaps with `EGL_EXT_device_drm`

[1](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_physical_device_drm-fn1)^.
Unlike the EGL extension, this extension does not expose a string
containing the name of the device file and instead exposes device minor
numbers.

DRM defines multiple device node types. Each physical device may have
one primary node and one render node associated. Physical devices may
have no primary node (e.g. if the device does not have a display
subsystem), may have no render node (e.g. if it is a software rendering
engine), or may have neither (e.g. if it is a software rendering engine
without a display subsystem).

To query DRM properties for a physical device, chain
[VkPhysicalDeviceDrmPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDrmPropertiesEXT.html)
to [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html).

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceDrmPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDrmPropertiesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_PHYSICAL_DEVICE_DRM_EXTENSION_NAME`

- `VK_EXT_PHYSICAL_DEVICE_DRM_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRM_PROPERTIES_EXT`

## <a href="#_references" class="anchor"></a>References

1.  <span id="VK_EXT_physical_device_drm-fn1"></span>
    [`EGL_EXT_device_drm`](https://registry.khronos.org/EGL/extensions/EXT/EGL_EXT_device_drm.txt)

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-06-09

  - First stable revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_physical_device_drm"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
