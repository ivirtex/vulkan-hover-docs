# VK_MSFT_layered_driver(3) Manual Page

## Name

VK_MSFT_layered_driver - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

531

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Jesse Natalie <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_MSFT_layered_driver%5D%20@jenatali%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_MSFT_layered_driver%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jenatali</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_MSFT_layered_driver](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_MSFT_layered_driver.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-06-21

**IP Status**  
No known IP claims.

**Contributors**  
- Jesse Natalie, Microsoft

## <a href="#_description" class="anchor"></a>Description

This extension adds new physical device properties to allow applications
and the Vulkan ICD loader to understand when a physical device is
implemented as a layered driver on top of another underlying API.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceLayeredDriverPropertiesMSFT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredDriverPropertiesMSFT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkLayeredDriverUnderlyingApiMSFT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayeredDriverUnderlyingApiMSFT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_MSFT_LAYERED_DRIVER_EXTENSION_NAME`

- `VK_MSFT_LAYERED_DRIVER_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_DRIVER_PROPERTIES_MSFT`

## <a href="#_examples" class="anchor"></a>Examples

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-06-21 (Jesse Natalie)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_MSFT_layered_driver"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
