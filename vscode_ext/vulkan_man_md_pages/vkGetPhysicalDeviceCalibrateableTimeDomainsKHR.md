# vkGetPhysicalDeviceCalibrateableTimeDomainsKHR(3) Manual Page

## Name

vkGetPhysicalDeviceCalibrateableTimeDomainsKHR - Query calibrateable
time domains



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the set of time domains for which a physical device supports
timestamp calibration, call:

``` c
// Provided by VK_KHR_calibrated_timestamps
VkResult vkGetPhysicalDeviceCalibrateableTimeDomainsKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pTimeDomainCount,
    VkTimeDomainKHR*                            pTimeDomains);
```

or the equivalent command

``` c
// Provided by VK_EXT_calibrated_timestamps
VkResult vkGetPhysicalDeviceCalibrateableTimeDomainsEXT(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pTimeDomainCount,
    VkTimeDomainKHR*                            pTimeDomains);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the set of
  calibrateable time domains.

- `pTimeDomainCount` is a pointer to an integer related to the number of
  calibrateable time domains available or queried, as described below.

- `pTimeDomains` is either `NULL` or a pointer to an array of
  [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimeDomainKHR.html) values, indicating the
  supported calibrateable time domains.

## <a href="#_description" class="anchor"></a>Description

If `pTimeDomains` is `NULL`, then the number of calibrateable time
domains supported for the given `physicalDevice` is returned in
`pTimeDomainCount`. Otherwise, `pTimeDomainCount` **must** point to a
variable set by the user to the number of elements in the `pTimeDomains`
array, and on return the variable is overwritten with the number of
values actually written to `pTimeDomains`. If the value of
`pTimeDomainCount` is less than the number of calibrateable time domains
supported, at most `pTimeDomainCount` values will be written to
`pTimeDomains`, and `VK_INCOMPLETE` will be returned instead of
`VK_SUCCESS`, to indicate that not all the available time domains were
returned.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-pTimeDomainCount-parameter"
  id="VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-pTimeDomainCount-parameter"></a>
  VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-pTimeDomainCount-parameter  
  `pTimeDomainCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-pTimeDomains-parameter"
  id="VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-pTimeDomains-parameter"></a>
  VUID-vkGetPhysicalDeviceCalibrateableTimeDomainsKHR-pTimeDomains-parameter  
  If the value referenced by `pTimeDomainCount` is not `0`, and
  `pTimeDomains` is not `NULL`, `pTimeDomains` **must** be a valid
  pointer to an array of `pTimeDomainCount`
  [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimeDomainKHR.html) values

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_calibrated_timestamps](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_calibrated_timestamps.html),
[VK_KHR_calibrated_timestamps](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_calibrated_timestamps.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimeDomainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceCalibrateableTimeDomainsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
