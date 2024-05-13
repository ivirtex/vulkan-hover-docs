# VK_EXT_custom_border_color(3) Manual Page

## Name

VK_EXT_custom_border_color - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

288

## <a href="#_revision" class="anchor"></a>Revision

12

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_special_uses" class="anchor"></a>Special Uses

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">OpenGL / ES support</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">D3D support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Liam Middlebrook <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_custom_border_color%5D%20@liam-middlebrook%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_custom_border_color%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>liam-middlebrook</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-04-16

**IP Status**  
No known IP claims.

**Contributors**  
- Joshua Ashton, Valve

- Hans-Kristian Arntzen, Valve

- Philip Rebohle, Valve

- Liam Middlebrook, NVIDIA

- Jeff Bolz, NVIDIA

- Tobias Hector, AMD

- Faith Ekstrand, Intel

- Spencer Fricke, Samsung Electronics

- Graeme Leese, Broadcom

- Jesse Hall, Google

- Jan-Harald Fredriksen, ARM

- Tom Olson, ARM

- Stuart Smith, Imagination Technologies

- Donald Scorgie, Imagination Technologies

- Alex Walters, Imagination Technologies

- Peter Quayle, Imagination Technologies

## <a href="#_description" class="anchor"></a>Description

This extension provides cross-vendor functionality to specify a custom
border color for use when the sampler address mode
`VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER` is used.

To create a sampler which uses a custom border color set
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`borderColor` to one
of:

- `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT`

- `VK_BORDER_COLOR_INT_CUSTOM_EXT`

When `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT` or
`VK_BORDER_COLOR_INT_CUSTOM_EXT` is used, applications must provide a
[VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html)
in the `pNext` chain for
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html).

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceCustomBorderColorFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCustomBorderColorFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceCustomBorderColorPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCustomBorderColorPropertiesEXT.html)

- Extending [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html):

  - [VkSamplerCustomBorderColorCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCustomBorderColorCreateInfoEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_CUSTOM_BORDER_COLOR_EXTENSION_NAME`

- `VK_EXT_CUSTOM_BORDER_COLOR_SPEC_VERSION`

- Extending [VkBorderColor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBorderColor.html):

  - `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT`

  - `VK_BORDER_COLOR_INT_CUSTOM_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUSTOM_BORDER_COLOR_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUSTOM_BORDER_COLOR_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_SAMPLER_CUSTOM_BORDER_COLOR_CREATE_INFO_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should VkClearColorValue be used for the border color value, or
should we have our own struct/union? Do we need to specify the type of
the input values for the components? This is more of a concern if
VkClearColorValue is used here because it provides a union of
float,int,uint types.

**RESOLVED**: Will reuse existing VkClearColorValue structure in order
to easily take advantage of float,int,uint borderColor types.

2\) For hardware which supports a limited number of border colors what
happens if that number is exceeded? Should this be handled by the driver
unbeknownst to the application? In Revision 1 we had solved this issue
using a new Object type, however that may have lead to additional system
resource consumption which would otherwise not be required.

**RESOLVED**: Added
`VkPhysicalDeviceCustomBorderColorPropertiesEXT`::`maxCustomBorderColorSamplers`
for tracking implementation-specific limit, and Valid Usage statement
handling overflow.

3\) Should this be supported for immutable samplers at all, or by a
feature bit? Some implementations may not be able to support custom
border colors on immutable samplers — is it worthwhile enabling this to
work on them for implementations that can support it, or forbidding it
entirely.

**RESOLVED**: Samplers created with a custom border color are forbidden
from being immutable. This resolves concerns for implementations where
the custom border color is an index to a LUT instead of being directly
embedded into sampler state.

4\) Should UINT and SINT (unsigned integer and signed integer) border
color types be separated or should they be combined into one generic INT
(integer) type?

**RESOLVED**: Separating these does not make much sense as the existing
fixed border color types do not have this distinction, and there is no
reason in hardware to do so. This separation would also create
unnecessary work and considerations for the application.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-10-10 (Joshua Ashton)

  - Internal revisions.

- Revision 2, 2019-10-11 (Liam Middlebrook)

  - Remove VkCustomBorderColor object and associated functions

  - Add issues concerning HW limitations for custom border color count

- Revision 3, 2019-10-12 (Joshua Ashton)

  - Re-expose the limits for the maximum number of unique border colors

  - Add extra details about border color tracking

  - Fix typos

- Revision 4, 2019-10-12 (Joshua Ashton)

  - Changed maxUniqueCustomBorderColors to a uint32_t from a
    VkDeviceSize

- Revision 5, 2019-10-14 (Liam Middlebrook)

  - Added features bit

- Revision 6, 2019-10-15 (Joshua Ashton)

  - Type-ize VK_BORDER_COLOR_CUSTOM

  - Fix const-ness on `pNext` of VkSamplerCustomBorderColorCreateInfoEXT

- Revision 7, 2019-11-26 (Liam Middlebrook)

  - Renamed maxUniqueCustomBorderColors to maxCustomBorderColors

- Revision 8, 2019-11-29 (Joshua Ashton)

  - Renamed borderColor member of
    VkSamplerCustomBorderColorCreateInfoEXT to customBorderColor

- Revision 9, 2020-02-19 (Joshua Ashton)

  - Renamed maxCustomBorderColors to maxCustomBorderColorSamplers

- Revision 10, 2020-02-21 (Joshua Ashton)

  - Added format to VkSamplerCustomBorderColorCreateInfoEXT and feature
    bit

- Revision 11, 2020-04-07 (Joshua Ashton)

  - Dropped UINT/SINT border color differences, consolidated types

- Revision 12, 2020-04-16 (Joshua Ashton)

  - Renamed VK_BORDER_COLOR_CUSTOM_FLOAT_EXT to
    VK_BORDER_COLOR_FLOAT_CUSTOM_EXT for consistency

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_custom_border_color"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
