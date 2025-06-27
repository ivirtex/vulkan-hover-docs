# VK_NV_low_latency(3) Manual Page

## Name

VK_NV_low_latency - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

311

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Charles Hansen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_low_latency%5D%20@cshansen%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_low_latency%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cshansen</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-02-10

**Contributors**  
- Charles Hansen, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds the
[VkQueryLowLatencySupportNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryLowLatencySupportNV.html) structure,
a structure used to query support for NVIDIA Reflex.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html):

  - [VkQueryLowLatencySupportNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryLowLatencySupportNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_LOW_LATENCY_EXTENSION_NAME`

- `VK_NV_LOW_LATENCY_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_QUERY_LOW_LATENCY_SUPPORT_NV`

## <a href="#_issues" class="anchor"></a>Issues

1\) Why does `VkQueryLowLatencySupportNV` have output parameters in an
input chain?

**RESOLVED**: We are stuck with this for legacy reasons - we are aware
this is bad behavior and this should not be used as a precedent for
future extensions.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-02-10 (Charles Hansen)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_low_latency"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
