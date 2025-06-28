# vkGetPhysicalDeviceCalibrateableTimeDomainsKHR(3) Manual Page

## Name

vkGetPhysicalDeviceCalibrateableTimeDomainsKHR - Query calibrateable time domains



## [](#_c_specification)C Specification

To query the set of time domains for which a physical device supports timestamp calibration, call:

```c++
// Provided by VK_KHR_calibrated_timestamps
VkResult vkGetPhysicalDeviceCalibrateableTimeDomainsKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pTimeDomainCount,
    VkTimeDomainKHR*                            pTimeDomains);
```

or the equivalent command

```c++
// Provided by VK_EXT_calibrated_timestamps
VkResult vkGetPhysicalDeviceCalibrateableTimeDomainsEXT(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pTimeDomainCount,
    VkTimeDomainKHR*                            pTimeDomains);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the set of calibrateable time domains.
- `pTimeDomainCount` is a pointer to an integer related to the number of calibrateable time domains available or queried, as described below.
- `pTimeDomains` is either `NULL` or a pointer to an array of [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimeDomainKHR.html) values, indicating the supported calibrateable time domains.

## [](#_description)Description

If `pTimeDomains` is `NULL`, then the number of calibrateable time domains supported for the given `physicalDevice` is returned in `pTimeDomainCount`. Otherwise, `pTimeDomainCount` **must** point to a variable set by the application to the number of elements in the `pTimeDomains` array, and on return the variable is overwritten with the number of values actually written to `pTimeDomains`. If the value of `pTimeDomainCount` is less than the number of calibrateable time domains supported, at most `pTimeDomainCount` values will be written to `pTimeDomains`, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available time domains were returned.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-pTimeDomainCount-parameter)VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-pTimeDomainCount-parameter  
  `pTimeDomainCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-pTimeDomains-parameter)VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-pTimeDomains-parameter  
  If the value referenced by `pTimeDomainCount` is not `0`, and `pTimeDomains` is not `NULL`, `pTimeDomains` **must** be a valid pointer to an array of `pTimeDomainCount` [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimeDomainKHR.html) values

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_EXT\_calibrated\_timestamps](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_calibrated_timestamps.html), [VK\_KHR\_calibrated\_timestamps](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_calibrated_timestamps.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimeDomainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceCalibrateableTimeDomainsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0