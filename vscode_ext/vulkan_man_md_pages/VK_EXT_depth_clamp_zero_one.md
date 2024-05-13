# VK_EXT_depth_clamp_zero_one(3) Manual Page

## Name

VK_EXT_depth_clamp_zero_one - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

422

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Graeme Leese <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_depth_clamp_zero_one%5D%20@gnl21%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_depth_clamp_zero_one%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>gnl21</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-07-29

**Contributors**  
- Graeme Leese, Broadcom

## <a href="#_description" class="anchor"></a>Description

This extension gives defined behavior to fragment depth values which end
up outside the conventional \[0, 1\] range. It can be used to ensure
portability in edge cases of features like depthBias. The particular
behavior is chosen to match OpenGL to aid porting or emulation.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceDepthClampZeroOneFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthClampZeroOneFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_DEPTH_CLAMP_ZERO_ONE_EXTENSION_NAME`

- `VK_EXT_DEPTH_CLAMP_ZERO_ONE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_CLAMP_ZERO_ONE_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-07-29 (Graeme Leese)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_depth_clamp_zero_one"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
