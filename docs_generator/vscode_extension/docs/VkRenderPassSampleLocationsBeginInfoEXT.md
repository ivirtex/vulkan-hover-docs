# VkRenderPassSampleLocationsBeginInfoEXT(3) Manual Page

## Name

VkRenderPassSampleLocationsBeginInfoEXT - Structure specifying sample locations to use for the layout transition of custom sample locations compatible depth/stencil attachments



## [](#_c_specification)C Specification

The image layout of the depth aspect of a depth/stencil attachment referring to an image created with `VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT` is dependent on the last sample locations used to render to the image subresource, thus preserving the contents of such depth/stencil attachments across subpass boundaries requires the application to specify these sample locations whenever a layout transition of the attachment **may** occur. This information **can** be provided by adding a `VkRenderPassSampleLocationsBeginInfoEXT` structure to the `pNext` chain of `VkRenderPassBeginInfo`.

The `VkRenderPassSampleLocationsBeginInfoEXT` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_EXT_sample_locations
typedef struct VkRenderPassSampleLocationsBeginInfoEXT {
    VkStructureType                          sType;
    const void*                              pNext;
    uint32_t                                 attachmentInitialSampleLocationsCount;
    const VkAttachmentSampleLocationsEXT*    pAttachmentInitialSampleLocations;
    uint32_t                                 postSubpassSampleLocationsCount;
    const VkSubpassSampleLocationsEXT*       pPostSubpassSampleLocations;
} VkRenderPassSampleLocationsBeginInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `attachmentInitialSampleLocationsCount` is the number of elements in the `pAttachmentInitialSampleLocations` array.
- `pAttachmentInitialSampleLocations` is a pointer to an array of `attachmentInitialSampleLocationsCount` [VkAttachmentSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentSampleLocationsEXT.html) structures specifying the attachment indices and their corresponding sample location state. Each element of `pAttachmentInitialSampleLocations` **can** specify the sample location state to use in the automatic layout transition performed to transition a depth/stencil attachment from the initial layout of the attachment to the image layout specified for the attachment in the first subpass using it.
- `postSubpassSampleLocationsCount` is the number of elements in the `pPostSubpassSampleLocations` array.
- `pPostSubpassSampleLocations` is a pointer to an array of `postSubpassSampleLocationsCount` [VkSubpassSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassSampleLocationsEXT.html) structures specifying the subpass indices and their corresponding sample location state. Each element of `pPostSubpassSampleLocations` **can** specify the sample location state to use in the automatic layout transition performed to transition the depth/stencil attachment used by the specified subpass to the image layout specified in a dependent subpass or to the final layout of the attachment in case the specified subpass is the last subpass using that attachment. In addition, if [VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html)::`variableSampleLocations` is `VK_FALSE`, each element of `pPostSubpassSampleLocations` **must** specify the sample location state that matches the sample locations used by all pipelines that will be bound to a command buffer during the specified subpass. If `variableSampleLocations` is `VK_TRUE`, the sample locations used for rasterization do not depend on `pPostSubpassSampleLocations`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkRenderPassSampleLocationsBeginInfoEXT-sType-sType)VUID-VkRenderPassSampleLocationsBeginInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_SAMPLE_LOCATIONS_BEGIN_INFO_EXT`
- [](#VUID-VkRenderPassSampleLocationsBeginInfoEXT-pAttachmentInitialSampleLocations-parameter)VUID-VkRenderPassSampleLocationsBeginInfoEXT-pAttachmentInitialSampleLocations-parameter  
  If `attachmentInitialSampleLocationsCount` is not `0`, `pAttachmentInitialSampleLocations` **must** be a valid pointer to an array of `attachmentInitialSampleLocationsCount` valid [VkAttachmentSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentSampleLocationsEXT.html) structures
- [](#VUID-VkRenderPassSampleLocationsBeginInfoEXT-pPostSubpassSampleLocations-parameter)VUID-VkRenderPassSampleLocationsBeginInfoEXT-pPostSubpassSampleLocations-parameter  
  If `postSubpassSampleLocationsCount` is not `0`, `pPostSubpassSampleLocations` **must** be a valid pointer to an array of `postSubpassSampleLocationsCount` valid [VkSubpassSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassSampleLocationsEXT.html) structures

## [](#_see_also)See Also

[VK\_EXT\_sample\_locations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sample_locations.html), [VkAttachmentSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentSampleLocationsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubpassSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassSampleLocationsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassSampleLocationsBeginInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0