# VK_ARM_render_pass_striped(3) Manual Page

## Name

VK_ARM_render_pass_striped - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

425

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

        
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Version 1.1](#versions-1.1)  
     and  
     [VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html)  
or  
[Version 1.3](#versions-1.3)  

## <a href="#_contact" class="anchor"></a>Contact

- Jan-Harald Fredriksen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ARM_render_pass_striped%5D%20@janharaldfredriksen-arm%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ARM_render_pass_striped%20extension*"
  target="_blank"
  rel="nofollow noopener"><em></em>janharaldfredriksen-arm</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_ARM_render_pass_striped](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_ARM_render_pass_striped.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension adds the ability to split a render pass instance into
stripes, and to get a notification when rendering has completed for each
stripe.

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkRenderPassStripeInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeInfoARM.html)

- Extending [VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferSubmitInfo.html):

  - [VkRenderPassStripeSubmitInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeSubmitInfoARM.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRenderPassStripedFeaturesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRenderPassStripedFeaturesARM.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceRenderPassStripedPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRenderPassStripedPropertiesARM.html)

- Extending [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html),
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html):

  - [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeBeginInfoARM.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_ARM_RENDER_PASS_STRIPED_EXTENSION_NAME`

- `VK_ARM_RENDER_PASS_STRIPED_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RENDER_PASS_STRIPED_FEATURES_ARM`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RENDER_PASS_STRIPED_PROPERTIES_ARM`

  - `VK_STRUCTURE_TYPE_RENDER_PASS_STRIPE_BEGIN_INFO_ARM`

  - `VK_STRUCTURE_TYPE_RENDER_PASS_STRIPE_INFO_ARM`

  - `VK_STRUCTURE_TYPE_RENDER_PASS_STRIPE_SUBMIT_INFO_ARM`

## <a href="#_examples" class="anchor"></a>Examples

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-11-21

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_ARM_render_pass_striped"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
