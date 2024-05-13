# vkSetHdrMetadataEXT(3) Manual Page

## Name

vkSetHdrMetadataEXT - Set Hdr metadata



## <a href="#_c_specification" class="anchor"></a>C Specification

To provide Hdr metadata to an implementation, call:

``` c
// Provided by VK_EXT_hdr_metadata
void vkSetHdrMetadataEXT(
    VkDevice                                    device,
    uint32_t                                    swapchainCount,
    const VkSwapchainKHR*                       pSwapchains,
    const VkHdrMetadataEXT*                     pMetadata);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device where the swapchain(s) were created.

- `swapchainCount` is the number of swapchains included in
  `pSwapchains`.

- `pSwapchains` is a pointer to an array of `swapchainCount`
  [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) handles.

- `pMetadata` is a pointer to an array of `swapchainCount`
  [VkHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHdrMetadataEXT.html) structures.

## <a href="#_description" class="anchor"></a>Description

The metadata will be applied to the specified
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) objects at the next
`vkQueuePresentKHR` call using that
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) object. The metadata will persist
until a subsequent `vkSetHdrMetadataEXT` changes it.

Valid Usage (Implicit)

- <a href="#VUID-vkSetHdrMetadataEXT-device-parameter"
  id="VUID-vkSetHdrMetadataEXT-device-parameter"></a>
  VUID-vkSetHdrMetadataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkSetHdrMetadataEXT-pSwapchains-parameter"
  id="VUID-vkSetHdrMetadataEXT-pSwapchains-parameter"></a>
  VUID-vkSetHdrMetadataEXT-pSwapchains-parameter  
  `pSwapchains` **must** be a valid pointer to an array of
  `swapchainCount` valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) handles

- <a href="#VUID-vkSetHdrMetadataEXT-pMetadata-parameter"
  id="VUID-vkSetHdrMetadataEXT-pMetadata-parameter"></a>
  VUID-vkSetHdrMetadataEXT-pMetadata-parameter  
  `pMetadata` **must** be a valid pointer to an array of
  `swapchainCount` valid [VkHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHdrMetadataEXT.html)
  structures

- <a href="#VUID-vkSetHdrMetadataEXT-swapchainCount-arraylength"
  id="VUID-vkSetHdrMetadataEXT-swapchainCount-arraylength"></a>
  VUID-vkSetHdrMetadataEXT-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

- <a href="#VUID-vkSetHdrMetadataEXT-pSwapchains-parent"
  id="VUID-vkSetHdrMetadataEXT-pSwapchains-parent"></a>
  VUID-vkSetHdrMetadataEXT-pSwapchains-parent  
  Each element of `pSwapchains` **must** have been created, allocated,
  or retrieved from `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_hdr_metadata](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_hdr_metadata.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHdrMetadataEXT.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkSetHdrMetadataEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
