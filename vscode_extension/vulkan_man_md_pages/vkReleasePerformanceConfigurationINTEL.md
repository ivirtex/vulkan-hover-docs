# vkReleasePerformanceConfigurationINTEL(3) Manual Page

## Name

vkReleasePerformanceConfigurationINTEL - Release a configuration to
capture performance data



## <a href="#_c_specification" class="anchor"></a>C Specification

To release a device performance configuration, call:

``` c
// Provided by VK_INTEL_performance_query
VkResult vkReleasePerformanceConfigurationINTEL(
    VkDevice                                    device,
    VkPerformanceConfigurationINTEL             configuration);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated to the configuration object to
  release.

- `configuration` is the configuration object to release.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkReleasePerformanceConfigurationINTEL-configuration-02737"
  id="VUID-vkReleasePerformanceConfigurationINTEL-configuration-02737"></a>
  VUID-vkReleasePerformanceConfigurationINTEL-configuration-02737  
  `configuration` **must** not be released before all command buffers
  submitted while the configuration was set are in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending state</a>

Valid Usage (Implicit)

- <a href="#VUID-vkReleasePerformanceConfigurationINTEL-device-parameter"
  id="VUID-vkReleasePerformanceConfigurationINTEL-device-parameter"></a>
  VUID-vkReleasePerformanceConfigurationINTEL-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkReleasePerformanceConfigurationINTEL-configuration-parameter"
  id="VUID-vkReleasePerformanceConfigurationINTEL-configuration-parameter"></a>
  VUID-vkReleasePerformanceConfigurationINTEL-configuration-parameter  
  If `configuration` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `configuration` **must** be a valid
  [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationINTEL.html)
  handle

- <a
  href="#VUID-vkReleasePerformanceConfigurationINTEL-configuration-parent"
  id="VUID-vkReleasePerformanceConfigurationINTEL-configuration-parent"></a>
  VUID-vkReleasePerformanceConfigurationINTEL-configuration-parent  
  If `configuration` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `configuration` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkReleasePerformanceConfigurationINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
