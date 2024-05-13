# vkGetValidationCacheDataEXT(3) Manual Page

## Name

vkGetValidationCacheDataEXT - Get the data store from a validation cache



## <a href="#_c_specification" class="anchor"></a>C Specification

Data **can** be retrieved from a validation cache object using the
command:

``` c
// Provided by VK_EXT_validation_cache
VkResult vkGetValidationCacheDataEXT(
    VkDevice                                    device,
    VkValidationCacheEXT                        validationCache,
    size_t*                                     pDataSize,
    void*                                       pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the validation cache.

- `validationCache` is the validation cache to retrieve data from.

- `pDataSize` is a pointer to a value related to the amount of data in
  the validation cache, as described below.

- `pData` is either `NULL` or a pointer to a buffer.

## <a href="#_description" class="anchor"></a>Description

If `pData` is `NULL`, then the maximum size of the data that **can** be
retrieved from the validation cache, in bytes, is returned in
`pDataSize`. Otherwise, `pDataSize` **must** point to a variable set by
the user to the size of the buffer, in bytes, pointed to by `pData`, and
on return the variable is overwritten with the amount of data actually
written to `pData`. If `pDataSize` is less than the maximum size that
**can** be retrieved by the validation cache, at most `pDataSize` bytes
will be written to `pData`, and `vkGetValidationCacheDataEXT` will
return `VK_INCOMPLETE` instead of `VK_SUCCESS`, to indicate that not all
of the validation cache was returned.

Any data written to `pData` is valid and **can** be provided as the
`pInitialData` member of the
[VkValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheCreateInfoEXT.html)
structure passed to `vkCreateValidationCacheEXT`.

Two calls to `vkGetValidationCacheDataEXT` with the same parameters
**must** retrieve the same data unless a command that modifies the
contents of the cache is called between them.

Applications **can** store the data retrieved from the validation cache,
and use these data, possibly in a future run of the application, to
populate new validation cache objects. The results of validation,
however, **may** depend on the vendor ID, device ID, driver version, and
other details of the device. To enable applications to detect when
previously retrieved data is incompatible with the device, the initial
bytes written to `pData` **must** be a header consisting of the
following members:

| Offset | Size           | Meaning                                                                                                                                                 |
|--------|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| 0      | 4              | length in bytes of the entire validation cache header written as a stream of bytes, with the least significant byte first                               |
| 4      | 4              | a [VkValidationCacheHeaderVersionEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheHeaderVersionEXT.html) value written as a stream of bytes, with the least significant byte first |
| 8      | `VK_UUID_SIZE` | a layer commit ID expressed as a UUID, which uniquely identifies the version of the validation layers used to generate these validation results         |

Table 1. Layout for validation cache header version
`VK_VALIDATION_CACHE_HEADER_VERSION_ONE_EXT`

The first four bytes encode the length of the entire validation cache
header, in bytes. This value includes all fields in the header including
the validation cache version field and the size of the length field.

The next four bytes encode the validation cache version, as described
for
[VkValidationCacheHeaderVersionEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheHeaderVersionEXT.html).
A consumer of the validation cache **should** use the cache version to
interpret the remainder of the cache header.

If `pDataSize` is less than what is necessary to store this header,
nothing will be written to `pData` and zero will be written to
`pDataSize`.

Valid Usage (Implicit)

- <a href="#VUID-vkGetValidationCacheDataEXT-device-parameter"
  id="VUID-vkGetValidationCacheDataEXT-device-parameter"></a>
  VUID-vkGetValidationCacheDataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetValidationCacheDataEXT-validationCache-parameter"
  id="VUID-vkGetValidationCacheDataEXT-validationCache-parameter"></a>
  VUID-vkGetValidationCacheDataEXT-validationCache-parameter  
  `validationCache` **must** be a valid
  [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html) handle

- <a href="#VUID-vkGetValidationCacheDataEXT-pDataSize-parameter"
  id="VUID-vkGetValidationCacheDataEXT-pDataSize-parameter"></a>
  VUID-vkGetValidationCacheDataEXT-pDataSize-parameter  
  `pDataSize` **must** be a valid pointer to a `size_t` value

- <a href="#VUID-vkGetValidationCacheDataEXT-pData-parameter"
  id="VUID-vkGetValidationCacheDataEXT-pData-parameter"></a>
  VUID-vkGetValidationCacheDataEXT-pData-parameter  
  If the value referenced by `pDataSize` is not `0`, and `pData` is not
  `NULL`, `pData` **must** be a valid pointer to an array of `pDataSize`
  bytes

- <a href="#VUID-vkGetValidationCacheDataEXT-validationCache-parent"
  id="VUID-vkGetValidationCacheDataEXT-validationCache-parent"></a>
  VUID-vkGetValidationCacheDataEXT-validationCache-parent  
  `validationCache` **must** have been created, allocated, or retrieved
  from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_validation_cache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_validation_cache.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetValidationCacheDataEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
