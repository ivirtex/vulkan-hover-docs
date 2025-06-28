# VK\_EXT\_non\_seamless\_cube\_map(3) Manual Page

## Name

VK\_EXT\_non\_seamless\_cube\_map - device extension



## [](#_registered_extension_number)Registered Extension Number

423

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_uses)Special Uses

- [D3D support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)
- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Georg Lehmann [\[GitHub\]DadSchoorse](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_non_seamless_cube_map%5D%20%40DadSchoorse%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_non_seamless_cube_map%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_non\_seamless\_cube\_map](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_non_seamless_cube_map.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-09-04

**IP Status**

No known IP claims.

**Contributors**

- Georg Lehmann

## [](#_description)Description

This extension provides functionality to disable [cube map edge handling](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-cubemapedge) on a per sampler level which matches the behavior of other graphics APIs.

This extension may be useful for building translation layers for those APIs or for porting applications that rely on this cube map behavior.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceNonSeamlessCubeMapFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceNonSeamlessCubeMapFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_NON_SEAMLESS_CUBE_MAP_EXTENSION_NAME`
- `VK_EXT_NON_SEAMLESS_CUBE_MAP_SPEC_VERSION`
- Extending [VkSamplerCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateFlagBits.html):
  
  - `VK_SAMPLER_CREATE_NON_SEAMLESS_CUBE_MAP_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_NON_SEAMLESS_CUBE_MAP_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2021-09-04 (Georg Lehmann)
  
  - First Version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_non_seamless_cube_map)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0