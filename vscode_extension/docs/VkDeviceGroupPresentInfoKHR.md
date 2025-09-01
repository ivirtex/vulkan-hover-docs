# VkDeviceGroupPresentInfoKHR(3) Manual Page

## Name

VkDeviceGroupPresentInfoKHR - Mode and mask controlling which physical devices' images are presented



## [](#_c_specification)C Specification

If the `pNext` chain of [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html) includes a `VkDeviceGroupPresentInfoKHR` structure, then that structure includes an array of device masks and a device group present mode.

The `VkDeviceGroupPresentInfoKHR` structure is defined as:

```c++
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_swapchain
typedef struct VkDeviceGroupPresentInfoKHR {
    VkStructureType                        sType;
    const void*                            pNext;
    uint32_t                               swapchainCount;
    const uint32_t*                        pDeviceMasks;
    VkDeviceGroupPresentModeFlagBitsKHR    mode;
} VkDeviceGroupPresentInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `swapchainCount` is zero or the number of elements in `pDeviceMasks`.
- `pDeviceMasks` is a pointer to an array of device masks, one for each element of [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html)::`pSwapchains`.
- `mode` is a [VkDeviceGroupPresentModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html) value specifying the device group present mode that will be used for this present.

## [](#_description)Description

If `mode` is `VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR`, then each element of `pDeviceMasks` selects which instance of the swapchain image is presented. Each element of `pDeviceMasks` **must** have exactly one bit set, and the corresponding physical device **must** have a presentation engine as reported by [VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentCapabilitiesKHR.html).

If `mode` is `VK_DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR`, then each element of `pDeviceMasks` selects which instance of the swapchain image is presented. Each element of `pDeviceMasks` **must** have exactly one bit set, and some physical device in the logical device **must** include that bit in its [VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentCapabilitiesKHR.html)::`presentMask`.

If `mode` is `VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR`, then each element of `pDeviceMasks` selects which instances of the swapchain image are component-wise summed and the sum of those images is presented. If the sum in any component is outside the representable range, the value of that component is undefined. Each element of `pDeviceMasks` **must** have a value for which all set bits are set in one of the elements of [VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentCapabilitiesKHR.html)::`presentMask`.

If `mode` is `VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR`, then each element of `pDeviceMasks` selects which instance(s) of the swapchain images are presented. For each bit set in each element of `pDeviceMasks`, the corresponding physical device **must** have a presentation engine as reported by [VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentCapabilitiesKHR.html).

If `VkDeviceGroupPresentInfoKHR` is not provided or `swapchainCount` is zero then the masks are considered to be `1`. If `VkDeviceGroupPresentInfoKHR` is not provided, `mode` is considered to be `VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR`.

Valid Usage

- [](#VUID-VkDeviceGroupPresentInfoKHR-swapchainCount-01297)VUID-VkDeviceGroupPresentInfoKHR-swapchainCount-01297  
  `swapchainCount` **must** equal `0` or [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html)::`swapchainCount`
- [](#VUID-VkDeviceGroupPresentInfoKHR-mode-01298)VUID-VkDeviceGroupPresentInfoKHR-mode-01298  
  If `mode` is `VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR`, then each element of `pDeviceMasks` **must** have exactly one bit set, and the corresponding element of [VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentCapabilitiesKHR.html)::`presentMask` **must** be non-zero
- [](#VUID-VkDeviceGroupPresentInfoKHR-mode-01299)VUID-VkDeviceGroupPresentInfoKHR-mode-01299  
  If `mode` is `VK_DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR`, then each element of `pDeviceMasks` **must** have exactly one bit set, and some physical device in the logical device **must** include that bit in its [VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentCapabilitiesKHR.html)::`presentMask`
- [](#VUID-VkDeviceGroupPresentInfoKHR-mode-01300)VUID-VkDeviceGroupPresentInfoKHR-mode-01300  
  If `mode` is `VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR`, then each element of `pDeviceMasks` **must** have a value for which all set bits are set in one of the elements of [VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentCapabilitiesKHR.html)::`presentMask`
- [](#VUID-VkDeviceGroupPresentInfoKHR-mode-01301)VUID-VkDeviceGroupPresentInfoKHR-mode-01301  
  If `mode` is `VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR`, then for each bit set in each element of `pDeviceMasks`, the corresponding element of [VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentCapabilitiesKHR.html)::`presentMask` **must** be non-zero
- [](#VUID-VkDeviceGroupPresentInfoKHR-pDeviceMasks-01302)VUID-VkDeviceGroupPresentInfoKHR-pDeviceMasks-01302  
  The value of each element of `pDeviceMasks` **must** be equal to the device mask passed in [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireNextImageInfoKHR.html)::`deviceMask` when the image index was last acquired
- [](#VUID-VkDeviceGroupPresentInfoKHR-mode-01303)VUID-VkDeviceGroupPresentInfoKHR-mode-01303  
  `mode` **must** have exactly one bit set, and that bit **must** have been included in [VkDeviceGroupSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupSwapchainCreateInfoKHR.html)::`modes`

Valid Usage (Implicit)

- [](#VUID-VkDeviceGroupPresentInfoKHR-sType-sType)VUID-VkDeviceGroupPresentInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR`
- [](#VUID-VkDeviceGroupPresentInfoKHR-pDeviceMasks-parameter)VUID-VkDeviceGroupPresentInfoKHR-pDeviceMasks-parameter  
  If `swapchainCount` is not `0`, `pDeviceMasks` **must** be a valid pointer to an array of `swapchainCount` `uint32_t` values
- [](#VUID-VkDeviceGroupPresentInfoKHR-mode-parameter)VUID-VkDeviceGroupPresentInfoKHR-mode-parameter  
  `mode` **must** be a valid [VkDeviceGroupPresentModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html) value

## [](#_see_also)See Also

[VK\_KHR\_device\_group](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_device_group.html), [VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkDeviceGroupPresentModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceGroupPresentInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0