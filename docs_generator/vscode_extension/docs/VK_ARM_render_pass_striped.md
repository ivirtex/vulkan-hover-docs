# VK\_ARM\_render\_pass\_striped(3) Manual Page

## Name

VK\_ARM\_render\_pass\_striped - device extension



## [](#_registered_extension_number)Registered Extension Number

425

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Vulkan Version 1.1](#versions-1.1)  
     and  
     [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_contact)Contact

- Jan-Harald Fredriksen [\[GitHub\]janharaldfredriksen-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ARM_render_pass_striped%5D%20%40janharaldfredriksen-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ARM_render_pass_striped%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_ARM\_render\_pass\_striped](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_ARM_render_pass_striped.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-11-21

**IP Status**

No known IP claims.

**Contributors**

- Jan-Harald Fredriksen, Arm
- Lisa Wu, Arm
- Torbjorn Nilsson, Arm
- Ying-Chieh Chen, Mediatek
- Jim Chiu, Mediatek

## [](#_description)Description

This extension adds the ability to split a render pass instance into stripes, and to get a notification when rendering has completed for each stripe.

## [](#_new_structures)New Structures

- [VkRenderPassStripeInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeInfoARM.html)
- Extending [VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferSubmitInfo.html):
  
  - [VkRenderPassStripeSubmitInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeSubmitInfoARM.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRenderPassStripedFeaturesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRenderPassStripedFeaturesARM.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceRenderPassStripedPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRenderPassStripedPropertiesARM.html)
- Extending [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html), [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html):
  
  - [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeBeginInfoARM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_ARM_RENDER_PASS_STRIPED_EXTENSION_NAME`
- `VK_ARM_RENDER_PASS_STRIPED_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RENDER_PASS_STRIPED_FEATURES_ARM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RENDER_PASS_STRIPED_PROPERTIES_ARM`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_STRIPE_BEGIN_INFO_ARM`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_STRIPE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_STRIPE_SUBMIT_INFO_ARM`

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2023-11-21
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_ARM_render_pass_striped)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0