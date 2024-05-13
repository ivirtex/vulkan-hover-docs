# VK_EXT_frame_boundary(3) Manual Page

## Name

VK_EXT_frame_boundary - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

376

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- James Fitzpatrick <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_frame_boundary%5D%20@jamesfitzpatrick%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_frame_boundary%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jamesfitzpatrick</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_frame_boundary](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_frame_boundary.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-06-14

**Contributors**  
- James Fitzpatrick, Imagination Technologies

- Hugues Evrard, Google

- Melih Yasin Yalcin, Google

- Andrew Garrard, Imagination Technologies

- Jan-Harald Fredriksen, Arm

- Vassili Nikolaev, NVIDIA

- Ting Wei, Huawei

## <a href="#_description" class="anchor"></a>Description

[VK_EXT_frame_boundary](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_frame_boundary.html) is a device
extension that helps **tools** (such as debuggers) to group queue
submissions per frames in non-trivial scenarios, typically when
[vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html) is not a relevant frame
boundary delimiter.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceFrameBoundaryFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFrameBoundaryFeaturesEXT.html)

- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html),
  [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo2.html),
  [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html),
  [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindSparseInfo.html):

  - [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkFrameBoundaryFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryFlagBitsEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkFrameBoundaryFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_FRAME_BOUNDARY_EXTENSION_NAME`

- `VK_EXT_FRAME_BOUNDARY_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_FRAME_BOUNDARY_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAME_BOUNDARY_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 0, 2022-01-14 (Hugues Evard)

  - Initial proposal

- Revision 1, 2023-06-14 (James Fitzpatrick)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_frame_boundary"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
