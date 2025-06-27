# VK_EXT_image_robustness(3) Manual Page

## Name

VK_EXT_image_robustness - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

336

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Graeme Leese <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_image_robustness%5D%20@gnl21%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_image_robustness%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>gnl21</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-04-27

**IP Status**  
No known IP claims.

**Contributors**  
- Graeme Leese, Broadcom

- Jan-Harald Fredriksen, ARM

- Jeff Bolz, NVIDIA

- Spencer Fricke, Samsung

- Courtney Goeltzenleuchter, Google

- Slawomir Cygan, Intel

## <a href="#_description" class="anchor"></a>Description

This extension adds stricter requirements for how out of bounds reads
from images are handled. Rather than returning undefined values, most
out of bounds reads return R, G, and B values of zero and alpha values
of either zero or one. Components not present in the image format may be
set to zero or to values based on the format as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-conversion-to-rgba"
target="_blank" rel="noopener">Conversion to RGBA</a>.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceImageRobustnessFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageRobustnessFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_IMAGE_ROBUSTNESS_EXTENSION_NAME`

- `VK_EXT_IMAGE_ROBUSTNESS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ROBUSTNESS_FEATURES_EXT`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Functionality in this extension is included in core Vulkan 1.3, with the
EXT suffix omitted. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_issues" class="anchor"></a>Issues

1.  How does this extension differ from VK_EXT_robustness2?

The guarantees provided by this extension are a subset of those provided
by the robustImageAccess2 feature of VK_EXT_robustness2. Where this
extension allows return values of (0, 0, 0, 0) or (0, 0, 0, 1),
robustImageAccess2 requires that a particular value dependent on the
image format be returned. This extension provides no guarantees about
the values returned for an access to an invalid Lod.

## <a href="#_examples" class="anchor"></a>Examples

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-04-27 (Graeme Leese)

- Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_robustness"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
