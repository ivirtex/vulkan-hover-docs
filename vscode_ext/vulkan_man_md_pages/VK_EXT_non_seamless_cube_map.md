# VK_EXT_non_seamless_cube_map(3) Manual Page

## Name

VK_EXT_non_seamless_cube_map - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

423

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_special_uses" class="anchor"></a>Special Uses

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">D3D support</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">OpenGL / ES support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Georg Lehmann <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_non_seamless_cube_map%5D%20@DadSchoorse%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_non_seamless_cube_map%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>DadSchoorse</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_non_seamless_cube_map](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_non_seamless_cube_map.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-09-04

**IP Status**  
No known IP claims.

**Contributors**  
- Georg Lehmann

## <a href="#_description" class="anchor"></a>Description

This extension provides functionality to disable <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-cubemapedge"
target="_blank" rel="noopener">cube map edge handling</a> on a per
sampler level which matches the behavior of other graphics APIs.

This extension may be useful for building translation layers for those
APIs or for porting applications that rely on this cube map behavior.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceNonSeamlessCubeMapFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceNonSeamlessCubeMapFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_NON_SEAMLESS_CUBE_MAP_EXTENSION_NAME`

- `VK_EXT_NON_SEAMLESS_CUBE_MAP_SPEC_VERSION`

- Extending [VkSamplerCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateFlagBits.html):

  - `VK_SAMPLER_CREATE_NON_SEAMLESS_CUBE_MAP_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_NON_SEAMLESS_CUBE_MAP_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-09-04 (Georg Lehmann)

  - First Version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_non_seamless_cube_map"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
