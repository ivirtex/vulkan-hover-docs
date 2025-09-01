# vkSetHdrMetadataEXT(3) Manual Page

## Name

vkSetHdrMetadataEXT - Set HDR metadata



## [](#_c_specification)C Specification

To provide HDR metadata to an implementation, call:

```c++
// Provided by VK_EXT_hdr_metadata
void vkSetHdrMetadataEXT(
    VkDevice                                    device,
    uint32_t                                    swapchainCount,
    const VkSwapchainKHR*                       pSwapchains,
    const VkHdrMetadataEXT*                     pMetadata);
```

## [](#_parameters)Parameters

- `device` is the logical device where the swapchain(s) were created.
- `swapchainCount` is the number of swapchains included in `pSwapchains`.
- `pSwapchains` is a pointer to an array of `swapchainCount` [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handles.
- `pMetadata` is a pointer to an array of `swapchainCount` [VkHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHdrMetadataEXT.html) structures.

## [](#_description)Description

The metadata will be applied to the specified [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) objects at the next [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) call using that [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) object. The metadata will persist until a subsequent `vkSetHdrMetadataEXT` changes it.

Valid Usage (Implicit)

- [](#VUID-vkSetHdrMetadataEXT-device-parameter)VUID-vkSetHdrMetadataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkSetHdrMetadataEXT-pSwapchains-parameter)VUID-vkSetHdrMetadataEXT-pSwapchains-parameter  
  `pSwapchains` **must** be a valid pointer to an array of `swapchainCount` valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handles
- [](#VUID-vkSetHdrMetadataEXT-pMetadata-parameter)VUID-vkSetHdrMetadataEXT-pMetadata-parameter  
  `pMetadata` **must** be a valid pointer to an array of `swapchainCount` valid [VkHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHdrMetadataEXT.html) structures
- [](#VUID-vkSetHdrMetadataEXT-swapchainCount-arraylength)VUID-vkSetHdrMetadataEXT-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`
- [](#VUID-vkSetHdrMetadataEXT-pSwapchains-parent)VUID-vkSetHdrMetadataEXT-pSwapchains-parent  
  Each element of `pSwapchains` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_EXT\_hdr\_metadata](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_hdr_metadata.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHdrMetadataEXT.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSetHdrMetadataEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0