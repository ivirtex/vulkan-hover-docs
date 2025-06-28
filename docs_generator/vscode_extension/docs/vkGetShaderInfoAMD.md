# vkGetShaderInfoAMD(3) Manual Page

## Name

vkGetShaderInfoAMD - Get information about a shader in a pipeline



## [](#_c_specification)C Specification

Information about a particular shader that has been compiled as part of a pipeline object can be extracted by calling:

```c++
// Provided by VK_AMD_shader_info
VkResult vkGetShaderInfoAMD(
    VkDevice                                    device,
    VkPipeline                                  pipeline,
    VkShaderStageFlagBits                       shaderStage,
    VkShaderInfoTypeAMD                         infoType,
    size_t*                                     pInfoSize,
    void*                                       pInfo);
```

## [](#_parameters)Parameters

- `device` is the device that created `pipeline`.
- `pipeline` is the target of the query.
- `shaderStage` is a [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) specifying the particular shader within the pipeline about which information is being queried.
- `infoType` describes what kind of information is being queried.
- `pInfoSize` is a pointer to a value related to the amount of data the query returns, as described below.
- `pInfo` is either `NULL` or a pointer to a buffer.

## [](#_description)Description

If `pInfo` is `NULL`, then the maximum size of the information that **can** be retrieved about the shader, in bytes, is returned in `pInfoSize`. Otherwise, `pInfoSize` **must** point to a variable set by the application to the size of the buffer, in bytes, pointed to by `pInfo`, and on return the variable is overwritten with the amount of data actually written to `pInfo`. If `pInfoSize` is less than the maximum size that **can** be retrieved by the pipeline cache, then at most `pInfoSize` bytes will be written to `pInfo`, and `VK_INCOMPLETE` will be returned, instead of `VK_SUCCESS`, to indicate that not all required of the pipeline cache was returned.

Not all information is available for every shader and implementations may not support all kinds of information for any shader. When a certain type of information is unavailable, the function returns `VK_ERROR_FEATURE_NOT_PRESENT`.

If information is successfully and fully queried, the function will return `VK_SUCCESS`.

For `infoType` `VK_SHADER_INFO_TYPE_STATISTICS_AMD`, a `VkShaderStatisticsInfoAMD` structure will be written to the buffer pointed to by `pInfo`. This structure will be populated with statistics regarding the physical device resources used by that shader along with other miscellaneous information and is described in further detail below.

For `infoType` `VK_SHADER_INFO_TYPE_DISASSEMBLY_AMD`, `pInfo` is a pointer to a null-terminated UTF-8 string containing human-readable disassembly. The exact formatting and contents of the disassembly string are vendor-specific.

The formatting and contents of all other types of information, including `infoType` `VK_SHADER_INFO_TYPE_BINARY_AMD`, are left to the vendor and are not further specified by this extension.

Note

This query does not behave consistently with the behavior described in [Opaque Binary Data Results](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-binaryresults), for historical reasons.

If the amount of data available is larger than the passed `pInfoSize`, the query returns up to the size of the passed buffer, and signals overflow with a `VK_INCOMPLETE` success status instead of returning a `VK_ERROR_NOT_ENOUGH_SPACE_KHR` error status.

Valid Usage (Implicit)

- [](#VUID-vkGetShaderInfoAMD-device-parameter)VUID-vkGetShaderInfoAMD-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetShaderInfoAMD-pipeline-parameter)VUID-vkGetShaderInfoAMD-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-vkGetShaderInfoAMD-shaderStage-parameter)VUID-vkGetShaderInfoAMD-shaderStage-parameter  
  `shaderStage` **must** be a valid [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) value
- [](#VUID-vkGetShaderInfoAMD-infoType-parameter)VUID-vkGetShaderInfoAMD-infoType-parameter  
  `infoType` **must** be a valid [VkShaderInfoTypeAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderInfoTypeAMD.html) value
- [](#VUID-vkGetShaderInfoAMD-pInfoSize-parameter)VUID-vkGetShaderInfoAMD-pInfoSize-parameter  
  `pInfoSize` **must** be a valid pointer to a `size_t` value
- [](#VUID-vkGetShaderInfoAMD-pInfo-parameter)VUID-vkGetShaderInfoAMD-pInfo-parameter  
  If the value referenced by `pInfoSize` is not `0`, and `pInfo` is not `NULL`, `pInfo` **must** be a valid pointer to an array of `pInfoSize` bytes
- [](#VUID-vkGetShaderInfoAMD-pipeline-parent)VUID-vkGetShaderInfoAMD-pipeline-parent  
  `pipeline` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_FEATURE_NOT_PRESENT`
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## [](#_see_also)See Also

[VK\_AMD\_shader\_info](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_shader_info.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkShaderInfoTypeAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderInfoTypeAMD.html), [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetShaderInfoAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0