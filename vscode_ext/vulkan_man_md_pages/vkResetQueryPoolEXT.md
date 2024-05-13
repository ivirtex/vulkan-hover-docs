# vkResetQueryPool(3) Manual Page

## Name

vkResetQueryPool - Reset queries in a query pool



## <a href="#_c_specification" class="anchor"></a>C Specification

To reset a range of queries in a query pool on the host, call:

``` c
// Provided by VK_VERSION_1_2
void vkResetQueryPool(
    VkDevice                                    device,
    VkQueryPool                                 queryPool,
    uint32_t                                    firstQuery,
    uint32_t                                    queryCount);
```

or the equivalent command

``` c
// Provided by VK_EXT_host_query_reset
void vkResetQueryPoolEXT(
    VkDevice                                    device,
    VkQueryPool                                 queryPool,
    uint32_t                                    firstQuery,
    uint32_t                                    queryCount);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the query pool.

- `queryPool` is the handle of the query pool managing the queries being
  reset.

- `firstQuery` is the initial query index to reset.

- `queryCount` is the number of queries to reset.

## <a href="#_description" class="anchor"></a>Description

This command sets the status of query indices \[`firstQuery`,
`firstQuery` + `queryCount` - 1\] to unavailable.

If `queryPool` is `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` this command
sets the status of query indices \[`firstQuery`, `firstQuery` +
`queryCount` - 1\] to unavailable for each pass.

Valid Usage

- <a href="#VUID-vkResetQueryPool-firstQuery-09436"
  id="VUID-vkResetQueryPool-firstQuery-09436"></a>
  VUID-vkResetQueryPool-firstQuery-09436  
  `firstQuery` **must** be less than the number of queries in
  `queryPool`

- <a href="#VUID-vkResetQueryPool-firstQuery-09437"
  id="VUID-vkResetQueryPool-firstQuery-09437"></a>
  VUID-vkResetQueryPool-firstQuery-09437  
  The sum of `firstQuery` and `queryCount` **must** be less than or
  equal to the number of queries in `queryPool`

<!-- -->

- <a href="#VUID-vkResetQueryPool-None-02665"
  id="VUID-vkResetQueryPool-None-02665"></a>
  VUID-vkResetQueryPool-None-02665  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-hostQueryReset"
  target="_blank" rel="noopener"><code>hostQueryReset</code></a> feature
  **must** be enabled

- <a href="#VUID-vkResetQueryPool-firstQuery-02741"
  id="VUID-vkResetQueryPool-firstQuery-02741"></a>
  VUID-vkResetQueryPool-firstQuery-02741  
  Submitted commands that refer to the range specified by `firstQuery`
  and `queryCount` in `queryPool` **must** have completed execution

- <a href="#VUID-vkResetQueryPool-firstQuery-02742"
  id="VUID-vkResetQueryPool-firstQuery-02742"></a>
  VUID-vkResetQueryPool-firstQuery-02742  
  The range of queries specified by `firstQuery` and `queryCount` in
  `queryPool` **must** not be in use by calls to
  [vkGetQueryPoolResults](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetQueryPoolResults.html) or
  `vkResetQueryPool` in other threads

Valid Usage (Implicit)

- <a href="#VUID-vkResetQueryPool-device-parameter"
  id="VUID-vkResetQueryPool-device-parameter"></a>
  VUID-vkResetQueryPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkResetQueryPool-queryPool-parameter"
  id="VUID-vkResetQueryPool-queryPool-parameter"></a>
  VUID-vkResetQueryPool-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a href="#VUID-vkResetQueryPool-queryPool-parent"
  id="VUID-vkResetQueryPool-queryPool-parent"></a>
  VUID-vkResetQueryPool-queryPool-parent  
  `queryPool` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_query_reset](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_query_reset.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkResetQueryPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
