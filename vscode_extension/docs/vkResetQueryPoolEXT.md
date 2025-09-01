# vkResetQueryPool(3) Manual Page

## Name

vkResetQueryPool - Reset queries in a query pool



## [](#_c_specification)C Specification

To reset a range of queries in a query pool on the host, call:

```c++
// Provided by VK_VERSION_1_2
void vkResetQueryPool(
    VkDevice                                    device,
    VkQueryPool                                 queryPool,
    uint32_t                                    firstQuery,
    uint32_t                                    queryCount);
```

or the equivalent command

```c++
// Provided by VK_EXT_host_query_reset
void vkResetQueryPoolEXT(
    VkDevice                                    device,
    VkQueryPool                                 queryPool,
    uint32_t                                    firstQuery,
    uint32_t                                    queryCount);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the query pool.
- `queryPool` is the handle of the query pool managing the queries being reset.
- `firstQuery` is the initial query index to reset.
- `queryCount` is the number of queries to reset.

## [](#_description)Description

This command sets the status of query indices \[`firstQuery`, `firstQuery` + `queryCount` - 1] to unavailable.

If `queryPool` is `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` this command sets the status of query indices \[`firstQuery`, `firstQuery` + `queryCount` - 1] to unavailable for each pass.

Valid Usage

- [](#VUID-vkResetQueryPool-firstQuery-09436)VUID-vkResetQueryPool-firstQuery-09436  
  `firstQuery` **must** be less than the number of queries in `queryPool`
- [](#VUID-vkResetQueryPool-firstQuery-09437)VUID-vkResetQueryPool-firstQuery-09437  
  The sum of `firstQuery` and `queryCount` **must** be less than or equal to the number of queries in `queryPool`

<!--THE END-->

- [](#VUID-vkResetQueryPool-None-02665)VUID-vkResetQueryPool-None-02665  
  The [`hostQueryReset`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-hostQueryReset) feature **must** be enabled
- [](#VUID-vkResetQueryPool-firstQuery-02741)VUID-vkResetQueryPool-firstQuery-02741  
  Submitted commands that refer to the range specified by `firstQuery` and `queryCount` in `queryPool` **must** have completed execution
- [](#VUID-vkResetQueryPool-firstQuery-02742)VUID-vkResetQueryPool-firstQuery-02742  
  The range of queries specified by `firstQuery` and `queryCount` in `queryPool` **must** not be in use by calls to [vkGetQueryPoolResults](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetQueryPoolResults.html) or `vkResetQueryPool` in other threads

Valid Usage (Implicit)

- [](#VUID-vkResetQueryPool-device-parameter)VUID-vkResetQueryPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkResetQueryPool-queryPool-parameter)VUID-vkResetQueryPool-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkResetQueryPool-queryPool-parent)VUID-vkResetQueryPool-queryPool-parent  
  `queryPool` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_EXT\_host\_query\_reset](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_query_reset.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkResetQueryPool).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0