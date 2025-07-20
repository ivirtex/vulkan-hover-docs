# vkBuildMicromapsEXT(3) Manual Page

## Name

vkBuildMicromapsEXT - Build a micromap on the host



## [](#_c_specification)C Specification

To build micromaps on the host, call:

```c++
// Provided by VK_EXT_opacity_micromap
VkResult vkBuildMicromapsEXT(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    uint32_t                                    infoCount,
    const VkMicromapBuildInfoEXT*               pInfos);
```

## [](#_parameters)Parameters

- `device` is the `VkDevice` for which the micromaps are being built.
- `deferredOperation` is an optional [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) to [request deferral](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#deferred-host-operations-requesting) for this command.
- `infoCount` is the number of micromaps to build. It specifies the number of the `pInfos` that **must** be provided.
- `pInfos` is a pointer to an array of `infoCount` [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html) structures defining the geometry used to build each micromap.

## [](#_description)Description

This command fulfills the same task as [vkCmdBuildMicromapsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildMicromapsEXT.html) but is executed by the host.

The `vkBuildMicromapsEXT` command provides the ability to initiate multiple micromaps builds, however there is no ordering or synchronization implied between any of the individual micromap builds.

Note

This means that there **cannot** be any memory aliasing between any micromap memories or scratch memories being used by any of the builds.

Valid Usage

- [](#VUID-vkBuildMicromapsEXT-pInfos-07461)VUID-vkBuildMicromapsEXT-pInfos-07461  
  For each `pInfos`\[i], `dstMicromap` **must** have been created with a value of [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateInfoEXT.html)::`size` greater than or equal to the memory size required by the build operation, as returned by [vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMicromapBuildSizesEXT.html) with `pBuildInfo` = `pInfos`\[i]
- [](#VUID-vkBuildMicromapsEXT-mode-07462)VUID-vkBuildMicromapsEXT-mode-07462  
  The `mode` member of each element of `pInfos` **must** be a valid [VkBuildMicromapModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildMicromapModeEXT.html) value
- [](#VUID-vkBuildMicromapsEXT-dstMicromap-07463)VUID-vkBuildMicromapsEXT-dstMicromap-07463  
  The `dstMicromap` member of any element of `pInfos` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) handle
- [](#VUID-vkBuildMicromapsEXT-pInfos-07464)VUID-vkBuildMicromapsEXT-pInfos-07464  
  For each element of `pInfos` its `type` member **must** match the value of [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateInfoEXT.html)::`type` when its `dstMicromap` was created
- [](#VUID-vkBuildMicromapsEXT-dstMicromap-07465)VUID-vkBuildMicromapsEXT-dstMicromap-07465  
  The range of memory backing the `dstMicromap` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `dstMicromap` member of any other element of `pInfos`, which is accessed by this command
- [](#VUID-vkBuildMicromapsEXT-dstMicromap-07466)VUID-vkBuildMicromapsEXT-dstMicromap-07466  
  The range of memory backing the `dstMicromap` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `scratchData` member of any element of `pInfos` (including the same element), which is accessed by this command
- [](#VUID-vkBuildMicromapsEXT-scratchData-07467)VUID-vkBuildMicromapsEXT-scratchData-07467  
  The range of memory backing the `scratchData` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `scratchData` member of any other element of `pInfos`, which is accessed by this command

<!--THE END-->

- [](#VUID-vkBuildMicromapsEXT-pInfos-07552)VUID-vkBuildMicromapsEXT-pInfos-07552  
  For each element of `pInfos`, the `buffer` used to create its `dstMicromap` member **must** be bound to host-visible device memory
- [](#VUID-vkBuildMicromapsEXT-pInfos-07553)VUID-vkBuildMicromapsEXT-pInfos-07553  
  For each element of `pInfos`, all referenced addresses of `pInfos`\[i].`data.hostAddress` **must** be valid host memory
- [](#VUID-vkBuildMicromapsEXT-pInfos-07554)VUID-vkBuildMicromapsEXT-pInfos-07554  
  For each element of `pInfos`, all referenced addresses of `pInfos`\[i].`triangleArray.hostAddress` **must** be valid host memory
- [](#VUID-vkBuildMicromapsEXT-micromapHostCommands-07555)VUID-vkBuildMicromapsEXT-micromapHostCommands-07555  
  The [`VkPhysicalDeviceOpacityMicromapFeaturesEXT`::`micromapHostCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-micromapHostCommands) feature **must** be enabled
- [](#VUID-vkBuildMicromapsEXT-pInfos-07556)VUID-vkBuildMicromapsEXT-pInfos-07556  
  If `pInfos`\[i].`mode` is `VK_BUILD_MICROMAP_MODE_BUILD_EXT`, all addresses between `pInfos`\[i].`scratchData.hostAddress` and `pInfos`\[i].`scratchData.hostAddress` + N - 1 **must** be valid host memory, where N is given by the `buildScratchSize` member of the [VkMicromapBuildSizesInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildSizesInfoEXT.html) structure returned from a call to [vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMicromapBuildSizesEXT.html) with an identical [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html) structure and primitive count
- [](#VUID-vkBuildMicromapsEXT-pInfos-07557)VUID-vkBuildMicromapsEXT-pInfos-07557  
  For each element of `pInfos`, the `buffer` used to create its `dstMicromap` member **must** be bound to memory that was not allocated with multiple instances

Valid Usage (Implicit)

- [](#VUID-vkBuildMicromapsEXT-device-parameter)VUID-vkBuildMicromapsEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkBuildMicromapsEXT-deferredOperation-parameter)VUID-vkBuildMicromapsEXT-deferredOperation-parameter  
  If `deferredOperation` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `deferredOperation` **must** be a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) handle
- [](#VUID-vkBuildMicromapsEXT-pInfos-parameter)VUID-vkBuildMicromapsEXT-pInfos-parameter  
  `pInfos` **must** be a valid pointer to an array of `infoCount` valid [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html) structures
- [](#VUID-vkBuildMicromapsEXT-infoCount-arraylength)VUID-vkBuildMicromapsEXT-infoCount-arraylength  
  `infoCount` **must** be greater than `0`
- [](#VUID-vkBuildMicromapsEXT-deferredOperation-parent)VUID-vkBuildMicromapsEXT-deferredOperation-parent  
  If `deferredOperation` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_OPERATION_DEFERRED_KHR`
- `VK_OPERATION_NOT_DEFERRED_KHR`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkBuildMicromapsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0