# vkSetPrivateData(3) Manual Page

## Name

vkSetPrivateData - Associate data with a Vulkan object



## <a href="#_c_specification" class="anchor"></a>C Specification

To store application-defined data in a slot associated with a Vulkan
object, call:

``` c
// Provided by VK_VERSION_1_3
VkResult vkSetPrivateData(
    VkDevice                                    device,
    VkObjectType                                objectType,
    uint64_t                                    objectHandle,
    VkPrivateDataSlot                           privateDataSlot,
    uint64_t                                    data);
```

or the equivalent command

``` c
// Provided by VK_EXT_private_data
VkResult vkSetPrivateDataEXT(
    VkDevice                                    device,
    VkObjectType                                objectType,
    uint64_t                                    objectHandle,
    VkPrivateDataSlot                           privateDataSlot,
    uint64_t                                    data);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device that created the object.

- `objectType` is a [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html) specifying the
  type of object to associate data with.

- `objectHandle` is a handle to the object to associate data with.

- `privateDataSlot` is a handle to a
  [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html) specifying location of
  private data storage.

- `data` is application-defined data to associate the object with. This
  data will be stored at `privateDataSlot`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkSetPrivateData-objectHandle-04016"
  id="VUID-vkSetPrivateData-objectHandle-04016"></a>
  VUID-vkSetPrivateData-objectHandle-04016  
  `objectHandle` **must** be `device` or a child of `device`

- <a href="#VUID-vkSetPrivateData-objectHandle-04017"
  id="VUID-vkSetPrivateData-objectHandle-04017"></a>
  VUID-vkSetPrivateData-objectHandle-04017  
  `objectHandle` **must** be a valid handle to an object of type
  `objectType`

Valid Usage (Implicit)

- <a href="#VUID-vkSetPrivateData-device-parameter"
  id="VUID-vkSetPrivateData-device-parameter"></a>
  VUID-vkSetPrivateData-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkSetPrivateData-objectType-parameter"
  id="VUID-vkSetPrivateData-objectType-parameter"></a>
  VUID-vkSetPrivateData-objectType-parameter  
  `objectType` **must** be a valid [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html)
  value

- <a href="#VUID-vkSetPrivateData-privateDataSlot-parameter"
  id="VUID-vkSetPrivateData-privateDataSlot-parameter"></a>
  VUID-vkSetPrivateData-privateDataSlot-parameter  
  `privateDataSlot` **must** be a valid
  [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html) handle

- <a href="#VUID-vkSetPrivateData-privateDataSlot-parent"
  id="VUID-vkSetPrivateData-privateDataSlot-parent"></a>
  VUID-vkSetPrivateData-privateDataSlot-parent  
  `privateDataSlot` **must** have been created, allocated, or retrieved
  from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_private_data](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_private_data.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html),
[VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkSetPrivateData"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
