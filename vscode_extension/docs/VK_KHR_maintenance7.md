# VK_KHR_maintenance7(3) Manual Page

## Name

VK_KHR_maintenance7 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

563

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Mike Blumenkrantz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance7%5D%20@zmike%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance7%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>zmike</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_maintenance7](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_maintenance7.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2024-01-30

**Interactions and External Dependencies**  
**Contributors**  
- Mike Blumenkrantz, Valve

- Hans-Kristian Arntzen, Valve

- Pan Gao, Huawei

- Tobias Hector, AMD

- Jon Leech, Khronos

- Daniel Story, Nintendo

- Shahbaz Youssefi, Google

- Yiwei Zhang, Google

- Matthew Netsch, Qualcomm

## <a href="#_description" class="anchor"></a>Description

[VK_KHR_maintenance7](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance7.html) adds a collection of
minor features, none of which would warrant an entire extension of their
own.

The proposed new features are as follows:

- Add a property query to determine if a framebuffer writes to depth or
  stencil aspect does not trigger a write access in the sibling aspect.
  For example, this allows sampling stencil aspect as a texture while
  rendering to the sibling depth attachment and vice-versa given
  appropriate image layouts.

- Add a way to query information regarding the underlying devices in
  environments where the Vulkan implementation is provided through
  layered implementations. For example, running on Mesa/Venus, driver ID
  is returned as `VK_DRIVER_ID_MESA_VENUS`, but it can be necessary to
  know what the real driver under the hood is. The new
  [VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html)
  structure can be used to gather information regarding layers
  underneath the top-level physical device.

- Promote `VK_RENDERING_CONTENTS_INLINE_BIT_EXT` and
  `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_EXT` to KHR

- Add a limit to report the maximum total count of dynamic uniform
  buffers and dynamic storage buffers that can be included in a pipeline
  layout.

- Require that for an unsigned integer query, the 32-bit result value
  **must** be equal to the 32 least significant bits of the equivalent
  64-bit result value.

- Add query for robust access support when using fragment shading rate
  attachments

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceMaintenance7FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance7FeaturesKHR.html)

- Extending
  [VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html):

  - [VkPhysicalDeviceLayeredApiVulkanPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiVulkanPropertiesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceLayeredApiPropertiesListKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiPropertiesListKHR.html)

  - [VkPhysicalDeviceMaintenance7PropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance7PropertiesKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkPhysicalDeviceLayeredApiKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_MAINTENANCE_7_EXTENSION_NAME`

- `VK_KHR_MAINTENANCE_7_SPEC_VERSION`

- Extending [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFlagBits.html):

  - `VK_RENDERING_CONTENTS_INLINE_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_API_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_API_PROPERTIES_LIST_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_API_VULKAN_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_7_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_7_PROPERTIES_KHR`

- Extending [VkSubpassContents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassContents.html):

  - `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_KHR`

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2024-01-30 (Jon Leech)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_maintenance7"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
