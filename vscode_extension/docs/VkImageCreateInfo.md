# VkImageCreateInfo(3) Manual Page

## Name

VkImageCreateInfo - Structure specifying the parameters of a newly
created image object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkImageCreateInfo {
    VkStructureType          sType;
    const void*              pNext;
    VkImageCreateFlags       flags;
    VkImageType              imageType;
    VkFormat                 format;
    VkExtent3D               extent;
    uint32_t                 mipLevels;
    uint32_t                 arrayLayers;
    VkSampleCountFlagBits    samples;
    VkImageTiling            tiling;
    VkImageUsageFlags        usage;
    VkSharingMode            sharingMode;
    uint32_t                 queueFamilyIndexCount;
    const uint32_t*          pQueueFamilyIndices;
    VkImageLayout            initialLayout;
} VkImageCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html) describing
  additional parameters of the image.

- `imageType` is a [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html) value specifying the
  basic dimensionality of the image. Layers in array textures do not
  count as a dimension for the purposes of the image type.

- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) describing the format and type
  of the texel blocks that will be contained in the image.

- `extent` is a [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html) describing the number of
  data elements in each dimension of the base level.

- `mipLevels` describes the number of levels of detail available for
  minified sampling of the image.

- `arrayLayers` is the number of layers in the image.

- `samples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html)
  value specifying the number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling"
  target="_blank" rel="noopener">samples per texel</a>.

- `tiling` is a [VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html) value specifying the
  tiling arrangement of the texel blocks in memory.

- `usage` is a bitmask of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) describing the
  intended usage of the image.

- `sharingMode` is a [VkSharingMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html) value
  specifying the sharing mode of the image when it will be accessed by
  multiple queue families.

- `queueFamilyIndexCount` is the number of entries in the
  `pQueueFamilyIndices` array.

- `pQueueFamilyIndices` is a pointer to an array of queue families that
  will access this image. It is ignored if `sharingMode` is not
  `VK_SHARING_MODE_CONCURRENT`.

- `initialLayout` is a [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value
  specifying the initial [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) of all
  image subresources of the image. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-layouts"
  target="_blank" rel="noopener">Image Layouts</a>.

## <a href="#_description" class="anchor"></a>Description

Images created with `tiling` equal to `VK_IMAGE_TILING_LINEAR` have
further restrictions on their limits and capabilities compared to images
created with `tiling` equal to `VK_IMAGE_TILING_OPTIMAL`. Creation of
images with tiling `VK_IMAGE_TILING_LINEAR` **may** not be supported
unless other parameters meet all of the constraints:

- `imageType` is `VK_IMAGE_TYPE_2D`

- `format` is not a depth/stencil format

- `mipLevels` is 1

- `arrayLayers` is 1

- `samples` is `VK_SAMPLE_COUNT_1_BIT`

- `usage` only includes `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` and/or
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT`

Images created with one of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
target="_blank" rel="noopener">formats that require a sampler
Y′C<sub>B</sub>C<sub>R</sub> conversion</a>, have further restrictions
on their limits and capabilities compared to images created with other
formats. Creation of images with a format requiring <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
target="_blank" rel="noopener">Y′C<sub>B</sub>C<sub>R</sub>
conversion</a> **may** not be supported unless other parameters meet all
of the constraints:

- `imageType` is `VK_IMAGE_TYPE_2D`

- `mipLevels` is 1

- `arrayLayers` is 1, unless the `ycbcrImageArrays` feature is enabled,
  or otherwise indicated by
  [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html)::`maxArrayLayers`,
  as returned by
  [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html)

- `samples` is `VK_SAMPLE_COUNT_1_BIT`

Implementations **may** support additional limits and capabilities
beyond those listed above.

To determine the set of valid `usage` bits for a given format, call
[vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html).

If the size of the resultant image would exceed `maxResourceSize`, then
[vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html) **must** fail and return
`VK_ERROR_OUT_OF_DEVICE_MEMORY`. This failure **may** occur even when
all image creation parameters satisfy their valid usage requirements.

If the implementation reports `VK_TRUE` in
[VkPhysicalDeviceHostImageCopyPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostImageCopyPropertiesEXT.html)::`identicalMemoryTypeRequirements`,
usage of `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` **must** not affect the
memory type requirements of the image as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#sparsememory-memory-requirements"
target="_blank" rel="noopener">Sparse Resource Memory Requirements</a>
and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-association"
target="_blank" rel="noopener">Resource Memory Association</a>.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>For images created without
<code>VK_IMAGE_CREATE_EXTENDED_USAGE_BIT</code> a <code>usage</code> bit
is valid if it is supported for the format the image is created
with.</p>
<p>For images created with
<code>VK_IMAGE_CREATE_EXTENDED_USAGE_BIT</code> a <code>usage</code> bit
is valid if it is supported for at least one of the formats a
<code>VkImageView</code> created from the image <strong>can</strong>
have (see <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views"
target="_blank" rel="noopener">Image Views</a> for more
detail).</p></td>
</tr>
</tbody>
</table>

Image Creation Limits

Valid values for some image creation parameters are limited by a
numerical upper bound or by inclusion in a bitset. For example,
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`arrayLayers` is limited by
`imageCreateMaxArrayLayers`, defined below; and
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`samples` is limited by
`imageCreateSampleCounts`, also defined below.

Several limiting values are defined below, as well as assisting values
from which the limiting values are derived. The limiting values are
referenced by the relevant valid usage statements of
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html).

- Let `uint64_t imageCreateDrmFormatModifiers[]` be the set of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-drm-format-modifier"
  target="_blank" rel="noopener">Linux DRM format modifiers</a> that the
  resultant image **may** have.

  - If `tiling` is not `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then
    `imageCreateDrmFormatModifiers` is empty.

  - If [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`pNext` contains
    [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html),
    then `imageCreateDrmFormatModifiers` contains exactly one modifier,
    [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)::`drmFormatModifier`.

  - If [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`pNext` contains
    [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html),
    then `imageCreateDrmFormatModifiers` contains the entire array
    [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)::`pDrmFormatModifiers`.

- Let `VkBool32 imageCreateMaybeLinear` indicate if the resultant image
  may be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-linear-resource"
  target="_blank" rel="noopener">linear</a>.

  - If `tiling` is `VK_IMAGE_TILING_LINEAR`, then
    `imageCreateMaybeLinear` is `VK_TRUE`.

  - If `tiling` is `VK_IMAGE_TILING_OPTIMAL`, then
    `imageCreateMaybeLinear` is `VK_FALSE`.

  - If `tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then
    `imageCreateMaybeLinear` is `VK_TRUE` if and only if
    `imageCreateDrmFormatModifiers` contains `DRM_FORMAT_MOD_LINEAR`.

- Let `VkFormatFeatureFlags imageCreateFormatFeatures` be the set of
  valid *format features* available during image creation.

  - If `tiling` is `VK_IMAGE_TILING_LINEAR`, then
    `imageCreateFormatFeatures` is the value of
    [VkFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html)::`linearTilingFeatures`
    found by calling
    [vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html)
    with parameter `format` equal to
    [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`format`.

  - If `tiling` is `VK_IMAGE_TILING_OPTIMAL`, and if the `pNext` chain
    includes no [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)
    or [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatQNX.html) structure with
    non-zero `externalFormat`, then `imageCreateFormatFeatures` is the
    value of
    [VkFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html)::`optimalTilingFeatures`
    found by calling
    [vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html)
    with parameter `format` equal to
    [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`format`.

  - If `tiling` is `VK_IMAGE_TILING_OPTIMAL`, and if the `pNext` chain
    includes a [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)
    structure with non-zero `externalFormat`, then
    `imageCreateFormatFeatures` is the value of
    [VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html)::`formatFeatures`
    obtained by
    [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html)
    with a matching `externalFormat` value.

  - If `tiling` is `VK_IMAGE_TILING_OPTIMAL`, and if the `pNext` chain
    includes a [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatQNX.html) structure
    with non-zero `externalFormat`, then `imageCreateFormatFeatures` is
    the value of
    [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferFormatPropertiesQNX.html)::`formatFeatures`
    obtained by
    [vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetScreenBufferPropertiesQNX.html)
    with a matching `externalFormat` value.

  - If the `pNext` chain includes a
    [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html)
    structure, then `imageCreateFormatFeatures` is the value of
    [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html)::`formatFeatures`
    found by calling
    [vkGetBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferCollectionPropertiesFUCHSIA.html)
    with a parameter `collection` equal to
    [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html)::`collection`

  - If `tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then the
    value of `imageCreateFormatFeatures` is found by calling
    [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html)
    with
    [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html)::`format`
    equal to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`format` and
    with
    [VkDrmFormatModifierPropertiesListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesListEXT.html)
    chained into [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html); by
    collecting all members of the returned array
    [VkDrmFormatModifierPropertiesListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesListEXT.html)::`pDrmFormatModifierProperties`
    whose `drmFormatModifier` belongs to
    `imageCreateDrmFormatModifiers`; and by taking the bitwise
    intersection, over the collected array members, of
    `drmFormatModifierTilingFeatures`. (The resultant
    `imageCreateFormatFeatures` **may** be empty).

- Let `VkImageFormatProperties2 imageCreateImageFormatPropertiesList[]`
  be defined as follows.

  - If [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`pNext` contains no
    [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html) or
    [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatQNX.html) structure with
    non-zero `externalFormat`, then
    `imageCreateImageFormatPropertiesList` is the list of structures
    obtained by calling
    [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html),
    possibly multiple times, as follows:

    - The parameters
      [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`format`,
      `imageType`, `tiling`, `usage`, and `flags` **must** be equal to
      those in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html).

    - If [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`pNext` contains a
      [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
      structure whose `handleTypes` is not `0`, then
      [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`pNext`
      **must** contain a
      [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html)
      structure whose `handleType` is not `0`; and
      [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
      **must** be called for each handle type in
      [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes`,
      successively setting
      [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html)::`handleType`
      on each call.

    - If [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`pNext` contains
      no
      [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
      structure, or contains a structure whose `handleTypes` is `0`,
      then
      [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`pNext`
      **must** either contain no
      [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html)
      structure, or contain a structure whose `handleType` is `0`.

    - If [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`pNext` contains a
      [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)
      structure then
      [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`pNext`
      **must** also contain the same
      [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)
      structure on each call.

    - If `tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then:

      - [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`pNext`
        **must** contain a
        [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html)
        structure where `sharingMode` is equal to
        [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`sharingMode`;

      - if `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, then
        `queueFamilyIndexCount` and `pQueueFamilyIndices` **must** be
        equal to those in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html);

      - if `flags` contains `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`, then
        the
        [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)
        structure included in the `pNext` chain of
        [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)
        **must** be equivalent to the one included in the `pNext` chain
        of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html);

      - if [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`pNext` contains
        a
        [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)
        structure, then the
        [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`pNext`
        chain **must** contain an equivalent structure;

      - [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
        **must** be called for each modifier in
        `imageCreateDrmFormatModifiers`, successively setting
        [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html)::`drmFormatModifier`
        on each call.

    - If `tiling` is not `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then
      [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`pNext`
      **must** contain no
      [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html)
      structure.

    - If any call to
      [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
      returns an error, then `imageCreateImageFormatPropertiesList` is
      defined to be the empty list.

  - If [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`pNext` contains a
    [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html) structure
    with non-zero `externalFormat`, then
    `imageCreateImageFormatPropertiesList` contains a single element
    where:

    - `VkImageFormatProperties`::`maxMipLevels` is
      ⌊log<sub>2</sub>(max(`extent.width`, `extent.height`,
      `extent.depth`))⌋ + 1.

    - `VkImageFormatProperties`::`maxArrayLayers` is
      [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxImageArrayLayers`.

    - Each component of `VkImageFormatProperties`::`maxExtent` is
      [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxImageDimension2D`.

    - `VkImageFormatProperties`::`sampleCounts` contains exactly
      `VK_SAMPLE_COUNT_1_BIT`.

- Let `uint32_t imageCreateMaxMipLevels` be the minimum value of
  [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html)::`maxMipLevels`
  in `imageCreateImageFormatPropertiesList`. The value is undefined if
  `imageCreateImageFormatPropertiesList` is empty.

- Let `uint32_t imageCreateMaxArrayLayers` be the minimum value of
  [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html)::`maxArrayLayers`
  in `imageCreateImageFormatPropertiesList`. The value is undefined if
  `imageCreateImageFormatPropertiesList` is empty.

- Let `VkExtent3D imageCreateMaxExtent` be the component-wise minimum
  over all
  [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html)::`maxExtent`
  values in `imageCreateImageFormatPropertiesList`. The value is
  undefined if `imageCreateImageFormatPropertiesList` is empty.

- Let `VkSampleCountFlags imageCreateSampleCounts` be the intersection
  of each
  [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html)::`sampleCounts`
  in `imageCreateImageFormatPropertiesList`. The value is undefined if
  `imageCreateImageFormatPropertiesList` is empty.

- Let `VkVideoFormatPropertiesKHR videoFormatProperties[]` be defined as
  follows.

  - If [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`pNext` contains a
    [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)
    structure, then `videoFormatProperties` is the list of structures
    obtained by calling
    [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html)
    with
    [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`imageUsage`
    equal to the `usage` member of
    [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) and
    [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`pNext`
    containing the same
    [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)
    structure chained to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html).

  - If [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`pNext` contains no
    [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)
    structure, then `videoFormatProperties` is an empty list.

- Let `VkBool32 supportedVideoFormat` indicate if the image parameters
  are supported by the specified video profiles.

  - `supportedVideoFormat` is `VK_TRUE` if there exists an element in
    the `videoFormatProperties` list for which all of the following
    conditions are true:

    - [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`format` equals
      [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)::`format`.

    - [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags` only contains
      bits also set in
      [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)::`imageCreateFlags`.

    - [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`imageType` equals
      [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)::`imageType`.

    - [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`tiling` equals
      [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)::`imageTiling`.

    - [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage` only contains
      bits also set in
      [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)::`imageUsageFlags`,
      or [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags` includes
      `VK_IMAGE_CREATE_EXTENDED_USAGE_BIT`.

  - Otherwise `supportedVideoFormat` is `VK_FALSE`.

Valid Usage

- <a href="#VUID-VkImageCreateInfo-imageCreateMaxMipLevels-02251"
  id="VUID-VkImageCreateInfo-imageCreateMaxMipLevels-02251"></a>
  VUID-VkImageCreateInfo-imageCreateMaxMipLevels-02251  
  Each of the following values (as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-creation-limits"
  target="_blank" rel="noopener">Image Creation Limits</a>) **must** not
  be undefined : `imageCreateMaxMipLevels`, `imageCreateMaxArrayLayers`,
  `imageCreateMaxExtent`, and `imageCreateSampleCounts`

- <a href="#VUID-VkImageCreateInfo-sharingMode-00941"
  id="VUID-VkImageCreateInfo-sharingMode-00941"></a>
  VUID-VkImageCreateInfo-sharingMode-00941  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`,
  `pQueueFamilyIndices` **must** be a valid pointer to an array of
  `queueFamilyIndexCount` `uint32_t` values

- <a href="#VUID-VkImageCreateInfo-sharingMode-00942"
  id="VUID-VkImageCreateInfo-sharingMode-00942"></a>
  VUID-VkImageCreateInfo-sharingMode-00942  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`,
  `queueFamilyIndexCount` **must** be greater than `1`

- <a href="#VUID-VkImageCreateInfo-sharingMode-01420"
  id="VUID-VkImageCreateInfo-sharingMode-01420"></a>
  VUID-VkImageCreateInfo-sharingMode-01420  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, each element of
  `pQueueFamilyIndices` **must** be unique and **must** be less than
  `pQueueFamilyPropertyCount` returned by either
  [vkGetPhysicalDeviceQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html)
  or
  [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html)
  for the `physicalDevice` that was used to create `device`

- <a href="#VUID-VkImageCreateInfo-pNext-01974"
  id="VUID-VkImageCreateInfo-pNext-01974"></a>
  VUID-VkImageCreateInfo-pNext-01974  
  If the `pNext` chain includes a
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html) structure, and
  its `externalFormat` member is non-zero the `format` **must** be
  `VK_FORMAT_UNDEFINED`

- <a href="#VUID-VkImageCreateInfo-pNext-01975"
  id="VUID-VkImageCreateInfo-pNext-01975"></a>
  VUID-VkImageCreateInfo-pNext-01975  
  If the `pNext` chain does not include a
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html) structure, or
  does and its `externalFormat` member is `0`, the `format` **must** not
  be `VK_FORMAT_UNDEFINED`

- <a href="#VUID-VkImageCreateInfo-extent-00944"
  id="VUID-VkImageCreateInfo-extent-00944"></a>
  VUID-VkImageCreateInfo-extent-00944  
  `extent.width` **must** be greater than `0`

- <a href="#VUID-VkImageCreateInfo-extent-00945"
  id="VUID-VkImageCreateInfo-extent-00945"></a>
  VUID-VkImageCreateInfo-extent-00945  
  `extent.height` **must** be greater than `0`

- <a href="#VUID-VkImageCreateInfo-extent-00946"
  id="VUID-VkImageCreateInfo-extent-00946"></a>
  VUID-VkImageCreateInfo-extent-00946  
  `extent.depth` **must** be greater than `0`

- <a href="#VUID-VkImageCreateInfo-mipLevels-00947"
  id="VUID-VkImageCreateInfo-mipLevels-00947"></a>
  VUID-VkImageCreateInfo-mipLevels-00947  
  `mipLevels` **must** be greater than `0`

- <a href="#VUID-VkImageCreateInfo-arrayLayers-00948"
  id="VUID-VkImageCreateInfo-arrayLayers-00948"></a>
  VUID-VkImageCreateInfo-arrayLayers-00948  
  `arrayLayers` **must** be greater than `0`

- <a href="#VUID-VkImageCreateInfo-flags-00949"
  id="VUID-VkImageCreateInfo-flags-00949"></a>
  VUID-VkImageCreateInfo-flags-00949  
  If `flags` contains `VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT`, `imageType`
  **must** be `VK_IMAGE_TYPE_2D`

- <a href="#VUID-VkImageCreateInfo-flags-08865"
  id="VUID-VkImageCreateInfo-flags-08865"></a>
  VUID-VkImageCreateInfo-flags-08865  
  If `flags` contains `VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT`,
  `extent.width` and `extent.height` **must** be equal

- <a href="#VUID-VkImageCreateInfo-flags-08866"
  id="VUID-VkImageCreateInfo-flags-08866"></a>
  VUID-VkImageCreateInfo-flags-08866  
  If `flags` contains `VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT`,
  `arrayLayers` **must** be greater than or equal to 6

- <a href="#VUID-VkImageCreateInfo-flags-02557"
  id="VUID-VkImageCreateInfo-flags-02557"></a>
  VUID-VkImageCreateInfo-flags-02557  
  If `flags` contains `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`,
  `imageType` **must** be `VK_IMAGE_TYPE_2D`

- <a href="#VUID-VkImageCreateInfo-flags-00950"
  id="VUID-VkImageCreateInfo-flags-00950"></a>
  VUID-VkImageCreateInfo-flags-00950  
  If `flags` contains `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT`,
  `imageType` **must** be `VK_IMAGE_TYPE_3D`

- <a href="#VUID-VkImageCreateInfo-flags-09403"
  id="VUID-VkImageCreateInfo-flags-09403"></a>
  VUID-VkImageCreateInfo-flags-09403  
  If `flags` contains `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT`, `flags`
  **must** not include `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT`,
  `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`, or
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`

- <a href="#VUID-VkImageCreateInfo-flags-07755"
  id="VUID-VkImageCreateInfo-flags-07755"></a>
  VUID-VkImageCreateInfo-flags-07755  
  If `flags` contains `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT`,
  `imageType` **must** be `VK_IMAGE_TYPE_3D`

- <a href="#VUID-VkImageCreateInfo-extent-02252"
  id="VUID-VkImageCreateInfo-extent-02252"></a>
  VUID-VkImageCreateInfo-extent-02252  
  `extent.width` **must** be less than or equal to
  `imageCreateMaxExtent.width` (as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-creation-limits"
  target="_blank" rel="noopener">Image Creation Limits</a>)

- <a href="#VUID-VkImageCreateInfo-extent-02253"
  id="VUID-VkImageCreateInfo-extent-02253"></a>
  VUID-VkImageCreateInfo-extent-02253  
  `extent.height` **must** be less than or equal to
  `imageCreateMaxExtent.height` (as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-creation-limits"
  target="_blank" rel="noopener">Image Creation Limits</a>)

- <a href="#VUID-VkImageCreateInfo-extent-02254"
  id="VUID-VkImageCreateInfo-extent-02254"></a>
  VUID-VkImageCreateInfo-extent-02254  
  `extent.depth` **must** be less than or equal to
  `imageCreateMaxExtent.depth` (as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-creation-limits"
  target="_blank" rel="noopener">Image Creation Limits</a>)

- <a href="#VUID-VkImageCreateInfo-imageType-00956"
  id="VUID-VkImageCreateInfo-imageType-00956"></a>
  VUID-VkImageCreateInfo-imageType-00956  
  If `imageType` is `VK_IMAGE_TYPE_1D`, both `extent.height` and
  `extent.depth` **must** be `1`

- <a href="#VUID-VkImageCreateInfo-imageType-00957"
  id="VUID-VkImageCreateInfo-imageType-00957"></a>
  VUID-VkImageCreateInfo-imageType-00957  
  If `imageType` is `VK_IMAGE_TYPE_2D`, `extent.depth` **must** be `1`

- <a href="#VUID-VkImageCreateInfo-mipLevels-00958"
  id="VUID-VkImageCreateInfo-mipLevels-00958"></a>
  VUID-VkImageCreateInfo-mipLevels-00958  
  `mipLevels` **must** be less than or equal to the number of levels in
  the complete mipmap chain based on `extent.width`, `extent.height`,
  and `extent.depth`

- <a href="#VUID-VkImageCreateInfo-mipLevels-02255"
  id="VUID-VkImageCreateInfo-mipLevels-02255"></a>
  VUID-VkImageCreateInfo-mipLevels-02255  
  `mipLevels` **must** be less than or equal to
  `imageCreateMaxMipLevels` (as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-creation-limits"
  target="_blank" rel="noopener">Image Creation Limits</a>)

- <a href="#VUID-VkImageCreateInfo-arrayLayers-02256"
  id="VUID-VkImageCreateInfo-arrayLayers-02256"></a>
  VUID-VkImageCreateInfo-arrayLayers-02256  
  `arrayLayers` **must** be less than or equal to
  `imageCreateMaxArrayLayers` (as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-creation-limits"
  target="_blank" rel="noopener">Image Creation Limits</a>)

- <a href="#VUID-VkImageCreateInfo-imageType-00961"
  id="VUID-VkImageCreateInfo-imageType-00961"></a>
  VUID-VkImageCreateInfo-imageType-00961  
  If `imageType` is `VK_IMAGE_TYPE_3D`, `arrayLayers` **must** be `1`

- <a href="#VUID-VkImageCreateInfo-samples-02257"
  id="VUID-VkImageCreateInfo-samples-02257"></a>
  VUID-VkImageCreateInfo-samples-02257  
  If `samples` is not `VK_SAMPLE_COUNT_1_BIT`, then `imageType` **must**
  be `VK_IMAGE_TYPE_2D`, `flags` **must** not contain
  `VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT`, `mipLevels` **must** be equal
  to `1`, and `imageCreateMaybeLinear` (as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-creation-limits"
  target="_blank" rel="noopener">Image Creation Limits</a>) **must** be
  `VK_FALSE`,

- <a href="#VUID-VkImageCreateInfo-samples-02558"
  id="VUID-VkImageCreateInfo-samples-02558"></a>
  VUID-VkImageCreateInfo-samples-02558  
  If `samples` is not `VK_SAMPLE_COUNT_1_BIT`, `usage` **must** not
  contain `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`

- <a href="#VUID-VkImageCreateInfo-usage-00963"
  id="VUID-VkImageCreateInfo-usage-00963"></a>
  VUID-VkImageCreateInfo-usage-00963  
  If `usage` includes `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`, then
  bits other than `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, and
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` **must** not be set

- <a href="#VUID-VkImageCreateInfo-usage-00964"
  id="VUID-VkImageCreateInfo-usage-00964"></a>
  VUID-VkImageCreateInfo-usage-00964  
  If `usage` includes `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`, or
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`, `extent.width` **must** be less
  than or equal to `VkPhysicalDeviceLimits`::`maxFramebufferWidth`

- <a href="#VUID-VkImageCreateInfo-usage-00965"
  id="VUID-VkImageCreateInfo-usage-00965"></a>
  VUID-VkImageCreateInfo-usage-00965  
  If `usage` includes `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`, or
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`, `extent.height` **must** be
  less than or equal to `VkPhysicalDeviceLimits`::`maxFramebufferHeight`

- <a href="#VUID-VkImageCreateInfo-fragmentDensityMapOffset-06514"
  id="VUID-VkImageCreateInfo-fragmentDensityMapOffset-06514"></a>
  VUID-VkImageCreateInfo-fragmentDensityMapOffset-06514  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentDensityMapOffset"
  target="_blank" rel="noopener"><code>fragmentDensityMapOffset</code></a>
  feature is not enabled and `usage` includes
  `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`, `extent.width` **must**
  be less than or equal to
  ⌈minFragmentDensityTexelSizewidth​maxFramebufferWidth​⌉

- <a href="#VUID-VkImageCreateInfo-fragmentDensityMapOffset-06515"
  id="VUID-VkImageCreateInfo-fragmentDensityMapOffset-06515"></a>
  VUID-VkImageCreateInfo-fragmentDensityMapOffset-06515  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentDensityMapOffset"
  target="_blank" rel="noopener"><code>fragmentDensityMapOffset</code></a>
  feature is not enabled and `usage` includes
  `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`, `extent.height`
  **must** be less than or equal to
  ⌈minFragmentDensityTexelSizeheight​maxFramebufferHeight​⌉

- <a href="#VUID-VkImageCreateInfo-usage-00966"
  id="VUID-VkImageCreateInfo-usage-00966"></a>
  VUID-VkImageCreateInfo-usage-00966  
  If `usage` includes `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`, `usage`
  **must** also contain at least one of
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, or
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-VkImageCreateInfo-samples-02258"
  id="VUID-VkImageCreateInfo-samples-02258"></a>
  VUID-VkImageCreateInfo-samples-02258  
  `samples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value that is set
  in `imageCreateSampleCounts` (as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-creation-limits"
  target="_blank" rel="noopener">Image Creation Limits</a>)

- <a href="#VUID-VkImageCreateInfo-usage-00968"
  id="VUID-VkImageCreateInfo-usage-00968"></a>
  VUID-VkImageCreateInfo-usage-00968  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderStorageImageMultisample"
  target="_blank"
  rel="noopener"><code>shaderStorageImageMultisample</code></a> feature
  is not enabled, and `usage` contains `VK_IMAGE_USAGE_STORAGE_BIT`,
  `samples` **must** be `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkImageCreateInfo-flags-00969"
  id="VUID-VkImageCreateInfo-flags-00969"></a>
  VUID-VkImageCreateInfo-flags-00969  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseBinding"
  target="_blank" rel="noopener"><code>sparseBinding</code></a> feature
  is not enabled, `flags` **must** not contain
  `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`

- <a href="#VUID-VkImageCreateInfo-flags-01924"
  id="VUID-VkImageCreateInfo-flags-01924"></a>
  VUID-VkImageCreateInfo-flags-01924  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseResidencyAliased"
  target="_blank" rel="noopener"><code>sparseResidencyAliased</code></a>
  feature is not enabled, `flags` **must** not contain
  `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT`

- <a href="#VUID-VkImageCreateInfo-tiling-04121"
  id="VUID-VkImageCreateInfo-tiling-04121"></a>
  VUID-VkImageCreateInfo-tiling-04121  
  If `tiling` is `VK_IMAGE_TILING_LINEAR`, `flags` **must** not contain
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`

- <a href="#VUID-VkImageCreateInfo-imageType-00970"
  id="VUID-VkImageCreateInfo-imageType-00970"></a>
  VUID-VkImageCreateInfo-imageType-00970  
  If `imageType` is `VK_IMAGE_TYPE_1D`, `flags` **must** not contain
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`

- <a href="#VUID-VkImageCreateInfo-imageType-00971"
  id="VUID-VkImageCreateInfo-imageType-00971"></a>
  VUID-VkImageCreateInfo-imageType-00971  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseResidencyImage2D"
  target="_blank" rel="noopener"><code>sparseResidencyImage2D</code></a>
  feature is not enabled, and `imageType` is `VK_IMAGE_TYPE_2D`, `flags`
  **must** not contain `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`

- <a href="#VUID-VkImageCreateInfo-imageType-00972"
  id="VUID-VkImageCreateInfo-imageType-00972"></a>
  VUID-VkImageCreateInfo-imageType-00972  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseResidencyImage3D"
  target="_blank" rel="noopener"><code>sparseResidencyImage3D</code></a>
  feature is not enabled, and `imageType` is `VK_IMAGE_TYPE_3D`, `flags`
  **must** not contain `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`

- <a href="#VUID-VkImageCreateInfo-imageType-00973"
  id="VUID-VkImageCreateInfo-imageType-00973"></a>
  VUID-VkImageCreateInfo-imageType-00973  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseResidency2Samples"
  target="_blank" rel="noopener"><code>sparseResidency2Samples</code></a>
  feature is not enabled, `imageType` is `VK_IMAGE_TYPE_2D`, and
  `samples` is `VK_SAMPLE_COUNT_2_BIT`, `flags` **must** not contain
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`

- <a href="#VUID-VkImageCreateInfo-imageType-00974"
  id="VUID-VkImageCreateInfo-imageType-00974"></a>
  VUID-VkImageCreateInfo-imageType-00974  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseResidency4Samples"
  target="_blank" rel="noopener"><code>sparseResidency4Samples</code></a>
  feature is not enabled, `imageType` is `VK_IMAGE_TYPE_2D`, and
  `samples` is `VK_SAMPLE_COUNT_4_BIT`, `flags` **must** not contain
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`

- <a href="#VUID-VkImageCreateInfo-imageType-00975"
  id="VUID-VkImageCreateInfo-imageType-00975"></a>
  VUID-VkImageCreateInfo-imageType-00975  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseResidency8Samples"
  target="_blank" rel="noopener"><code>sparseResidency8Samples</code></a>
  feature is not enabled, `imageType` is `VK_IMAGE_TYPE_2D`, and
  `samples` is `VK_SAMPLE_COUNT_8_BIT`, `flags` **must** not contain
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`

- <a href="#VUID-VkImageCreateInfo-imageType-00976"
  id="VUID-VkImageCreateInfo-imageType-00976"></a>
  VUID-VkImageCreateInfo-imageType-00976  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseResidency16Samples"
  target="_blank" rel="noopener"><code>sparseResidency16Samples</code></a>
  feature is not enabled, `imageType` is `VK_IMAGE_TYPE_2D`, and
  `samples` is `VK_SAMPLE_COUNT_16_BIT`, `flags` **must** not contain
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`

- <a href="#VUID-VkImageCreateInfo-flags-00987"
  id="VUID-VkImageCreateInfo-flags-00987"></a>
  VUID-VkImageCreateInfo-flags-00987  
  If `flags` contains `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` or
  `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT`, it **must** also contain
  `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`

- <a href="#VUID-VkImageCreateInfo-None-01925"
  id="VUID-VkImageCreateInfo-None-01925"></a>
  VUID-VkImageCreateInfo-None-01925  
  If any of the bits `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`,
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`, or
  `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT` are set,
  `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT` **must** not also be set

- <a href="#VUID-VkImageCreateInfo-flags-01890"
  id="VUID-VkImageCreateInfo-flags-01890"></a>
  VUID-VkImageCreateInfo-flags-01890  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-protectedMemory"
  target="_blank" rel="noopener"><code>protectedMemory</code></a>
  feature is not enabled, `flags` **must** not contain
  `VK_IMAGE_CREATE_PROTECTED_BIT`

- <a href="#VUID-VkImageCreateInfo-None-01891"
  id="VUID-VkImageCreateInfo-None-01891"></a>
  VUID-VkImageCreateInfo-None-01891  
  If any of the bits `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`,
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`, or
  `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT` are set,
  `VK_IMAGE_CREATE_PROTECTED_BIT` **must** not also be set

- <a href="#VUID-VkImageCreateInfo-pNext-00988"
  id="VUID-VkImageCreateInfo-pNext-00988"></a>
  VUID-VkImageCreateInfo-pNext-00988  
  If the `pNext` chain includes a
  [VkExternalMemoryImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfoNV.html)
  structure, it **must** not contain a
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
  structure

- <a href="#VUID-VkImageCreateInfo-pNext-00990"
  id="VUID-VkImageCreateInfo-pNext-00990"></a>
  VUID-VkImageCreateInfo-pNext-00990  
  If the `pNext` chain includes a
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
  structure, its `handleTypes` member **must** only contain bits that
  are also in
  [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatProperties.html)::`externalMemoryProperties.compatibleHandleTypes`,
  as returned by
  [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
  with `format`, `imageType`, `tiling`, `usage`, and `flags` equal to
  those in this structure, and with a
  [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html)
  structure included in the `pNext` chain, with a `handleType` equal to
  any one of the handle types specified in
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes`

- <a href="#VUID-VkImageCreateInfo-pNext-00991"
  id="VUID-VkImageCreateInfo-pNext-00991"></a>
  VUID-VkImageCreateInfo-pNext-00991  
  If the `pNext` chain includes a
  [VkExternalMemoryImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfoNV.html)
  structure, its `handleTypes` member **must** only contain bits that
  are also in
  [VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatPropertiesNV.html)::`externalMemoryProperties.compatibleHandleTypes`,
  as returned by
  [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)
  with `format`, `imageType`, `tiling`, `usage`, and `flags` equal to
  those in this structure, and with `externalHandleType` equal to any
  one of the handle types specified in
  [VkExternalMemoryImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfoNV.html)::`handleTypes`

- <a href="#VUID-VkImageCreateInfo-physicalDeviceCount-01421"
  id="VUID-VkImageCreateInfo-physicalDeviceCount-01421"></a>
  VUID-VkImageCreateInfo-physicalDeviceCount-01421  
  If the logical device was created with
  [VkDeviceGroupDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupDeviceCreateInfo.html)::`physicalDeviceCount`
  equal to 1, `flags` **must** not contain
  `VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT`

- <a href="#VUID-VkImageCreateInfo-flags-02259"
  id="VUID-VkImageCreateInfo-flags-02259"></a>
  VUID-VkImageCreateInfo-flags-02259  
  If `flags` contains `VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT`,
  then `mipLevels` **must** be one, `arrayLayers` **must** be one,
  `imageType` **must** be `VK_IMAGE_TYPE_2D`. and
  `imageCreateMaybeLinear` (as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-creation-limits"
  target="_blank" rel="noopener">Image Creation Limits</a>) **must** be
  `VK_FALSE`

- <a href="#VUID-VkImageCreateInfo-flags-01572"
  id="VUID-VkImageCreateInfo-flags-01572"></a>
  VUID-VkImageCreateInfo-flags-01572  
  If `flags` contains `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT`,
  then `format` **must** be a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#compressed_image_formats"
  target="_blank" rel="noopener">compressed image format</a>

- <a href="#VUID-VkImageCreateInfo-flags-01573"
  id="VUID-VkImageCreateInfo-flags-01573"></a>
  VUID-VkImageCreateInfo-flags-01573  
  If `flags` contains `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT`,
  then `flags` **must** also contain
  `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`

- <a href="#VUID-VkImageCreateInfo-initialLayout-00993"
  id="VUID-VkImageCreateInfo-initialLayout-00993"></a>
  VUID-VkImageCreateInfo-initialLayout-00993  
  `initialLayout` **must** be `VK_IMAGE_LAYOUT_UNDEFINED` or
  `VK_IMAGE_LAYOUT_PREINITIALIZED`

- <a href="#VUID-VkImageCreateInfo-pNext-01443"
  id="VUID-VkImageCreateInfo-pNext-01443"></a>
  VUID-VkImageCreateInfo-pNext-01443  
  If the `pNext` chain includes a
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
  or `VkExternalMemoryImageCreateInfoNV` structure whose `handleTypes`
  member is not `0`, `initialLayout` **must** be
  `VK_IMAGE_LAYOUT_UNDEFINED`

- <a href="#VUID-VkImageCreateInfo-format-06410"
  id="VUID-VkImageCreateInfo-format-06410"></a>
  VUID-VkImageCreateInfo-format-06410  
  If the image `format` is one of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">formats that require a sampler
  Y′C<sub>B</sub>C<sub>R</sub> conversion</a>, `mipLevels` **must** be 1

- <a href="#VUID-VkImageCreateInfo-format-06411"
  id="VUID-VkImageCreateInfo-format-06411"></a>
  VUID-VkImageCreateInfo-format-06411  
  If the image `format` is one of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">formats that require a sampler
  Y′C<sub>B</sub>C<sub>R</sub> conversion</a>, `samples` **must** be
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkImageCreateInfo-format-06412"
  id="VUID-VkImageCreateInfo-format-06412"></a>
  VUID-VkImageCreateInfo-format-06412  
  If the image `format` is one of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">formats that require a sampler
  Y′C<sub>B</sub>C<sub>R</sub> conversion</a>, `imageType` **must** be
  `VK_IMAGE_TYPE_2D`

- <a href="#VUID-VkImageCreateInfo-imageCreateFormatFeatures-02260"
  id="VUID-VkImageCreateInfo-imageCreateFormatFeatures-02260"></a>
  VUID-VkImageCreateInfo-imageCreateFormatFeatures-02260  
  If `format` is a *multi-planar* format, and if
  `imageCreateFormatFeatures` (as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-creation-limits"
  target="_blank" rel="noopener">Image Creation Limits</a>) does not
  contain `VK_FORMAT_FEATURE_DISJOINT_BIT`, then `flags` **must** not
  contain `VK_IMAGE_CREATE_DISJOINT_BIT`

- <a href="#VUID-VkImageCreateInfo-format-01577"
  id="VUID-VkImageCreateInfo-format-01577"></a>
  VUID-VkImageCreateInfo-format-01577  
  If `format` is not a *multi-planar* format, and `flags` does not
  include `VK_IMAGE_CREATE_ALIAS_BIT`, `flags` **must** not contain
  `VK_IMAGE_CREATE_DISJOINT_BIT`

- <a href="#VUID-VkImageCreateInfo-format-04712"
  id="VUID-VkImageCreateInfo-format-04712"></a>
  VUID-VkImageCreateInfo-format-04712  
  If `format` has a `_422` or `_420` suffix, `extent.width` **must** be
  a multiple of 2

- <a href="#VUID-VkImageCreateInfo-format-04713"
  id="VUID-VkImageCreateInfo-format-04713"></a>
  VUID-VkImageCreateInfo-format-04713  
  If `format` has a `_420` suffix, `extent.height` **must** be a
  multiple of 2

- <a href="#VUID-VkImageCreateInfo-format-09583"
  id="VUID-VkImageCreateInfo-format-09583"></a>
  VUID-VkImageCreateInfo-format-09583  
  If `format` is one of the `VK_FORMAT_PVTRC1_*_IMG` formats,
  `extent.width` **must** be a power of 2

- <a href="#VUID-VkImageCreateInfo-format-09584"
  id="VUID-VkImageCreateInfo-format-09584"></a>
  VUID-VkImageCreateInfo-format-09584  
  If `format` is one of the `VK_FORMAT_PVTRC1_*_IMG` formats,
  `extent.height` **must** be a power of 2

- <a href="#VUID-VkImageCreateInfo-tiling-02261"
  id="VUID-VkImageCreateInfo-tiling-02261"></a>
  VUID-VkImageCreateInfo-tiling-02261  
  If `tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then the
  `pNext` chain **must** include exactly one of
  [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)
  or
  [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)
  structures

- <a href="#VUID-VkImageCreateInfo-pNext-02262"
  id="VUID-VkImageCreateInfo-pNext-02262"></a>
  VUID-VkImageCreateInfo-pNext-02262  
  If the `pNext` chain includes a
  [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)
  or
  [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)
  structure, then `tiling` **must** be
  `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`

- <a href="#VUID-VkImageCreateInfo-tiling-02353"
  id="VUID-VkImageCreateInfo-tiling-02353"></a>
  VUID-VkImageCreateInfo-tiling-02353  
  If `tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT` and `flags`
  contains `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`, then the `pNext` chain
  **must** include a
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)
  structure with non-zero `viewFormatCount`

- <a href="#VUID-VkImageCreateInfo-flags-01533"
  id="VUID-VkImageCreateInfo-flags-01533"></a>
  VUID-VkImageCreateInfo-flags-01533  
  If `flags` contains
  `VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT` `format`
  **must** be a depth or depth/stencil format

- <a href="#VUID-VkImageCreateInfo-pNext-02393"
  id="VUID-VkImageCreateInfo-pNext-02393"></a>
  VUID-VkImageCreateInfo-pNext-02393  
  If the `pNext` chain includes a
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
  structure whose `handleTypes` member includes
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`,
  `imageType` **must** be `VK_IMAGE_TYPE_2D`

- <a href="#VUID-VkImageCreateInfo-pNext-02394"
  id="VUID-VkImageCreateInfo-pNext-02394"></a>
  VUID-VkImageCreateInfo-pNext-02394  
  If the `pNext` chain includes a
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
  structure whose `handleTypes` member includes
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`,
  `mipLevels` **must** either be `1` or equal to the number of levels in
  the complete mipmap chain based on `extent.width`, `extent.height`,
  and `extent.depth`

- <a href="#VUID-VkImageCreateInfo-pNext-02396"
  id="VUID-VkImageCreateInfo-pNext-02396"></a>
  VUID-VkImageCreateInfo-pNext-02396  
  If the `pNext` chain includes a
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html) structure
  whose `externalFormat` member is not `0`, `flags` **must** not include
  `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`

- <a href="#VUID-VkImageCreateInfo-pNext-02397"
  id="VUID-VkImageCreateInfo-pNext-02397"></a>
  VUID-VkImageCreateInfo-pNext-02397  
  If the `pNext` chain includes a
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html) structure
  whose `externalFormat` member is not `0`, `usage` **must** not include
  any usages except `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`, or `VK_IMAGE_USAGE_SAMPLED_BIT`

- <a href="#VUID-VkImageCreateInfo-pNext-09457"
  id="VUID-VkImageCreateInfo-pNext-09457"></a>
  VUID-VkImageCreateInfo-pNext-09457  
  If the `pNext` chain includes a
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html) structure
  whose `externalFormat` member is not `0`, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is not enabled, `usage` **must** not include
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` or
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-VkImageCreateInfo-pNext-02398"
  id="VUID-VkImageCreateInfo-pNext-02398"></a>
  VUID-VkImageCreateInfo-pNext-02398  
  If the `pNext` chain includes a
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html) structure
  whose `externalFormat` member is not `0`, `tiling` **must** be
  `VK_IMAGE_TILING_OPTIMAL`

- <a href="#VUID-VkImageCreateInfo-pNext-08951"
  id="VUID-VkImageCreateInfo-pNext-08951"></a>
  VUID-VkImageCreateInfo-pNext-08951  
  If the `pNext` chain includes a
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
  structure whose `handleTypes` member includes
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_SCREEN_BUFFER_BIT_QNX`, `imageType`
  **must** be `VK_IMAGE_TYPE_2D`

- <a href="#VUID-VkImageCreateInfo-pNext-08952"
  id="VUID-VkImageCreateInfo-pNext-08952"></a>
  VUID-VkImageCreateInfo-pNext-08952  
  If the `pNext` chain includes a
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
  structure whose `handleTypes` member includes
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_SCREEN_BUFFER_BIT_QNX`, `mipLevels`
  **must** either be `1` or equal to the number of levels in the
  complete mipmap chain based on `extent.width`, `extent.height`, and
  `extent.depth`

- <a href="#VUID-VkImageCreateInfo-pNext-08953"
  id="VUID-VkImageCreateInfo-pNext-08953"></a>
  VUID-VkImageCreateInfo-pNext-08953  
  If the `pNext` chain includes a
  [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatQNX.html) structure whose
  `externalFormat` member is not `0`, `flags` **must** not include
  `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`

- <a href="#VUID-VkImageCreateInfo-pNext-08954"
  id="VUID-VkImageCreateInfo-pNext-08954"></a>
  VUID-VkImageCreateInfo-pNext-08954  
  If the `pNext` chain includes a
  [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatQNX.html) structure whose
  `externalFormat` member is not `0`, `usage` **must** not include any
  usages except `VK_IMAGE_USAGE_SAMPLED_BIT`

- <a href="#VUID-VkImageCreateInfo-pNext-08955"
  id="VUID-VkImageCreateInfo-pNext-08955"></a>
  VUID-VkImageCreateInfo-pNext-08955  
  If the `pNext` chain includes a
  [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatQNX.html) structure whose
  `externalFormat` member is not `0`, `tiling` **must** be
  `VK_IMAGE_TILING_OPTIMAL`

- <a href="#VUID-VkImageCreateInfo-format-02795"
  id="VUID-VkImageCreateInfo-format-02795"></a>
  VUID-VkImageCreateInfo-format-02795  
  If `format` is a depth-stencil format, `usage` includes
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, and the `pNext` chain
  includes a
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)
  structure, then its
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`
  member **must** also include
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageCreateInfo-format-02796"
  id="VUID-VkImageCreateInfo-format-02796"></a>
  VUID-VkImageCreateInfo-format-02796  
  If `format` is a depth-stencil format, `usage` does not include
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, and the `pNext` chain
  includes a
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)
  structure, then its
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`
  member **must** also not include
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageCreateInfo-format-02797"
  id="VUID-VkImageCreateInfo-format-02797"></a>
  VUID-VkImageCreateInfo-format-02797  
  If `format` is a depth-stencil format, `usage` includes
  `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`, and the `pNext` chain
  includes a
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)
  structure, then its
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`
  member **must** also include `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`

- <a href="#VUID-VkImageCreateInfo-format-02798"
  id="VUID-VkImageCreateInfo-format-02798"></a>
  VUID-VkImageCreateInfo-format-02798  
  If `format` is a depth-stencil format, `usage` does not include
  `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`, and the `pNext` chain
  includes a
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)
  structure, then its
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`
  member **must** also not include
  `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`

- <a href="#VUID-VkImageCreateInfo-Format-02536"
  id="VUID-VkImageCreateInfo-Format-02536"></a>
  VUID-VkImageCreateInfo-Format-02536  
  If `Format` is a depth-stencil format and the `pNext` chain includes a
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)
  structure with its `stencilUsage` member including
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`, `extent.width` **must** be less
  than or equal to `VkPhysicalDeviceLimits`::`maxFramebufferWidth`

- <a href="#VUID-VkImageCreateInfo-format-02537"
  id="VUID-VkImageCreateInfo-format-02537"></a>
  VUID-VkImageCreateInfo-format-02537  
  If `format` is a depth-stencil format and the `pNext` chain includes a
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)
  structure with its `stencilUsage` member including
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`, `extent.height` **must** be
  less than or equal to `VkPhysicalDeviceLimits`::`maxFramebufferHeight`

- <a href="#VUID-VkImageCreateInfo-format-02538"
  id="VUID-VkImageCreateInfo-format-02538"></a>
  VUID-VkImageCreateInfo-format-02538  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderStorageImageMultisample"
  target="_blank"
  rel="noopener"><code>shaderStorageImageMultisample</code></a> feature
  is not enabled, `format` is a depth-stencil format and the `pNext`
  chain includes a
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)
  structure with its `stencilUsage` including
  `VK_IMAGE_USAGE_STORAGE_BIT`, `samples` **must** be
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkImageCreateInfo-flags-02050"
  id="VUID-VkImageCreateInfo-flags-02050"></a>
  VUID-VkImageCreateInfo-flags-02050  
  If `flags` contains `VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV`,
  `imageType` **must** be `VK_IMAGE_TYPE_2D` or `VK_IMAGE_TYPE_3D`

- <a href="#VUID-VkImageCreateInfo-flags-02051"
  id="VUID-VkImageCreateInfo-flags-02051"></a>
  VUID-VkImageCreateInfo-flags-02051  
  If `flags` contains `VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV`, it
  **must** not contain `VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT` and the
  `format` **must** not be a depth/stencil format

- <a href="#VUID-VkImageCreateInfo-flags-02052"
  id="VUID-VkImageCreateInfo-flags-02052"></a>
  VUID-VkImageCreateInfo-flags-02052  
  If `flags` contains `VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV` and
  `imageType` is `VK_IMAGE_TYPE_2D`, `extent.width` and `extent.height`
  **must** be greater than `1`

- <a href="#VUID-VkImageCreateInfo-flags-02053"
  id="VUID-VkImageCreateInfo-flags-02053"></a>
  VUID-VkImageCreateInfo-flags-02053  
  If `flags` contains `VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV` and
  `imageType` is `VK_IMAGE_TYPE_3D`, `extent.width`, `extent.height`,
  and `extent.depth` **must** be greater than `1`

- <a href="#VUID-VkImageCreateInfo-imageType-02082"
  id="VUID-VkImageCreateInfo-imageType-02082"></a>
  VUID-VkImageCreateInfo-imageType-02082  
  If `usage` includes
  `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`, `imageType`
  **must** be `VK_IMAGE_TYPE_2D`

- <a href="#VUID-VkImageCreateInfo-samples-02083"
  id="VUID-VkImageCreateInfo-samples-02083"></a>
  VUID-VkImageCreateInfo-samples-02083  
  If `usage` includes
  `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`, `samples`
  **must** be `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkImageCreateInfo-shadingRateImage-07727"
  id="VUID-VkImageCreateInfo-shadingRateImage-07727"></a>
  VUID-VkImageCreateInfo-shadingRateImage-07727  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shadingRateImage"
  target="_blank" rel="noopener"><code>shadingRateImage</code></a>
  feature is enabled and `usage` includes
  `VK_IMAGE_USAGE_SHADING_RATE_IMAGE_BIT_NV`, `tiling` **must** be
  `VK_IMAGE_TILING_OPTIMAL`

- <a href="#VUID-VkImageCreateInfo-flags-02565"
  id="VUID-VkImageCreateInfo-flags-02565"></a>
  VUID-VkImageCreateInfo-flags-02565  
  If `flags` contains `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`, `tiling`
  **must** be `VK_IMAGE_TILING_OPTIMAL`

- <a href="#VUID-VkImageCreateInfo-flags-02566"
  id="VUID-VkImageCreateInfo-flags-02566"></a>
  VUID-VkImageCreateInfo-flags-02566  
  If `flags` contains `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`, `imageType`
  **must** be `VK_IMAGE_TYPE_2D`

- <a href="#VUID-VkImageCreateInfo-flags-02567"
  id="VUID-VkImageCreateInfo-flags-02567"></a>
  VUID-VkImageCreateInfo-flags-02567  
  If `flags` contains `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`, `flags`
  **must** not contain `VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT`

- <a href="#VUID-VkImageCreateInfo-flags-02568"
  id="VUID-VkImageCreateInfo-flags-02568"></a>
  VUID-VkImageCreateInfo-flags-02568  
  If `flags` contains `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`, `mipLevels`
  **must** be `1`

- <a href="#VUID-VkImageCreateInfo-usage-04992"
  id="VUID-VkImageCreateInfo-usage-04992"></a>
  VUID-VkImageCreateInfo-usage-04992  
  If `usage` includes `VK_IMAGE_USAGE_INVOCATION_MASK_BIT_HUAWEI`,
  `tiling` **must** be `VK_IMAGE_TILING_LINEAR`

- <a href="#VUID-VkImageCreateInfo-imageView2DOn3DImage-04459"
  id="VUID-VkImageCreateInfo-imageView2DOn3DImage-04459"></a>
  VUID-VkImageCreateInfo-imageView2DOn3DImage-04459  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`imageView2DOn3DImage`
  is `VK_FALSE`, `flags` **must** not contain
  `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT`

- <a href="#VUID-VkImageCreateInfo-multisampleArrayImage-04460"
  id="VUID-VkImageCreateInfo-multisampleArrayImage-04460"></a>
  VUID-VkImageCreateInfo-multisampleArrayImage-04460  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`multisampleArrayImage`
  is `VK_FALSE`, and `samples` is not `VK_SAMPLE_COUNT_1_BIT`, then
  `arrayLayers` **must** be `1`

- <a href="#VUID-VkImageCreateInfo-pNext-06722"
  id="VUID-VkImageCreateInfo-pNext-06722"></a>
  VUID-VkImageCreateInfo-pNext-06722  
  If a [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)
  structure was included in the `pNext` chain and
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)::`viewFormatCount`
  is not zero, then each format in
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)::`pViewFormats`
  **must** either be compatible with the `format` as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatibility"
  target="_blank" rel="noopener">compatibility table</a> or, if `flags`
  contains `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT`, be an
  uncompressed format that is size-compatible with `format`

- <a href="#VUID-VkImageCreateInfo-flags-04738"
  id="VUID-VkImageCreateInfo-flags-04738"></a>
  VUID-VkImageCreateInfo-flags-04738  
  If `flags` does not contain `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT` and
  the `pNext` chain includes a
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)
  structure, then
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)::`viewFormatCount`
  **must** be `0` or `1`

- <a href="#VUID-VkImageCreateInfo-usage-04815"
  id="VUID-VkImageCreateInfo-usage-04815"></a>
  VUID-VkImageCreateInfo-usage-04815  
  If `usage` includes `VK_IMAGE_USAGE_VIDEO_DECODE_SRC_BIT_KHR`,
  `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`, or
  `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`, and `flags` does not
  include `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`, then the
  `pNext` chain **must** include a
  [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html) structure
  with `profileCount` greater than `0` and `pProfiles` including at
  least one [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)
  structure with a `videoCodecOperation` member specifying a decode
  operation

- <a href="#VUID-VkImageCreateInfo-usage-04816"
  id="VUID-VkImageCreateInfo-usage-04816"></a>
  VUID-VkImageCreateInfo-usage-04816  
  If `usage` includes `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`,
  `VK_IMAGE_USAGE_VIDEO_ENCODE_DST_BIT_KHR`, or
  `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`, and `flags` does not
  include `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`, then the
  `pNext` chain **must** include a
  [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html) structure
  with `profileCount` greater than `0` and `pProfiles` including at
  least one [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)
  structure with a `videoCodecOperation` member specifying an encode
  operation

- <a href="#VUID-VkImageCreateInfo-flags-08328"
  id="VUID-VkImageCreateInfo-flags-08328"></a>
  VUID-VkImageCreateInfo-flags-08328  
  If `flags` includes
  `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`, then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-videoMaintenance1"
  target="_blank" rel="noopener"><code>videoMaintenance1</code></a>
  **must** be enabled

- <a href="#VUID-VkImageCreateInfo-flags-08329"
  id="VUID-VkImageCreateInfo-flags-08329"></a>
  VUID-VkImageCreateInfo-flags-08329  
  If `flags` includes
  `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR` and `usage` does
  not include `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`, then `usage`
  **must** not include `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`

- <a href="#VUID-VkImageCreateInfo-flags-08331"
  id="VUID-VkImageCreateInfo-flags-08331"></a>
  VUID-VkImageCreateInfo-flags-08331  
  If `flags` includes
  `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`, then `usage`
  **must** not include `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`

- <a href="#VUID-VkImageCreateInfo-pNext-06811"
  id="VUID-VkImageCreateInfo-pNext-06811"></a>
  VUID-VkImageCreateInfo-pNext-06811  
  If the `pNext` chain includes a
  [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html) structure
  with `profileCount` greater than `0`, then `supportedVideoFormat`
  **must** be `VK_TRUE`

- <a href="#VUID-VkImageCreateInfo-pNext-06390"
  id="VUID-VkImageCreateInfo-pNext-06390"></a>
  VUID-VkImageCreateInfo-pNext-06390  
  If the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) is to be used to import memory from a
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html), a
  [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html)
  structure **must** be chained to `pNext`

- <a
  href="#VUID-VkImageCreateInfo-multisampledRenderToSingleSampled-06882"
  id="VUID-VkImageCreateInfo-multisampledRenderToSingleSampled-06882"></a>
  VUID-VkImageCreateInfo-multisampledRenderToSingleSampled-06882  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multisampledRenderToSingleSampled"
  target="_blank"
  rel="noopener"><code>multisampledRenderToSingleSampled</code></a>
  feature is not enabled, `flags` **must** not contain
  `VK_IMAGE_CREATE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_BIT_EXT`

- <a href="#VUID-VkImageCreateInfo-flags-06883"
  id="VUID-VkImageCreateInfo-flags-06883"></a>
  VUID-VkImageCreateInfo-flags-06883  
  If `flags` contains
  `VK_IMAGE_CREATE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_BIT_EXT`,
  `samples` **must** be `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkImageCreateInfo-pNext-06743"
  id="VUID-VkImageCreateInfo-pNext-06743"></a>
  VUID-VkImageCreateInfo-pNext-06743  
  If the `pNext` chain includes a
  [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)
  structure, `format` is a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">multi-planar</a> format, and
  [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)::`flags`
  includes `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT`, then
  [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)::`compressionControlPlaneCount`
  **must** be equal to the number of planes in `format`

- <a href="#VUID-VkImageCreateInfo-pNext-06744"
  id="VUID-VkImageCreateInfo-pNext-06744"></a>
  VUID-VkImageCreateInfo-pNext-06744  
  If the `pNext` chain includes a
  [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)
  structure, `format` is not a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">multi-planar</a> format, and
  [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)::`flags`
  includes `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT`, then
  [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)::`compressionControlPlaneCount`
  **must** be 1

- <a href="#VUID-VkImageCreateInfo-pNext-06746"
  id="VUID-VkImageCreateInfo-pNext-06746"></a>
  VUID-VkImageCreateInfo-pNext-06746  
  If the `pNext` chain includes a
  [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)
  structure, it **must** not contain a
  [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)
  structure

- <a href="#VUID-VkImageCreateInfo-flags-08104"
  id="VUID-VkImageCreateInfo-flags-08104"></a>
  VUID-VkImageCreateInfo-flags-08104  
  If `flags` includes
  `VK_IMAGE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBufferCaptureReplay"
  target="_blank"
  rel="noopener"><code>descriptorBufferCaptureReplay</code></a> feature
  **must** be enabled

- <a href="#VUID-VkImageCreateInfo-pNext-08105"
  id="VUID-VkImageCreateInfo-pNext-08105"></a>
  VUID-VkImageCreateInfo-pNext-08105  
  If the `pNext` chain includes a
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
  structure, `flags` **must** contain
  `VK_IMAGE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`

- <a href="#VUID-VkImageCreateInfo-pNext-06783"
  id="VUID-VkImageCreateInfo-pNext-06783"></a>
  VUID-VkImageCreateInfo-pNext-06783  
  If the `pNext` chain includes a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure, its `exportObjectType` member **must** be either
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT` or
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_IOSURFACE_BIT_EXT`

- <a href="#VUID-VkImageCreateInfo-pNext-06784"
  id="VUID-VkImageCreateInfo-pNext-06784"></a>
  VUID-VkImageCreateInfo-pNext-06784  
  If the `pNext` chain includes a
  [VkImportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalTextureInfoEXT.html)
  structure its `plane` member **must** be
  `VK_IMAGE_ASPECT_PLANE_0_BIT`, `VK_IMAGE_ASPECT_PLANE_1_BIT`, or
  `VK_IMAGE_ASPECT_PLANE_2_BIT`

- <a href="#VUID-VkImageCreateInfo-pNext-06785"
  id="VUID-VkImageCreateInfo-pNext-06785"></a>
  VUID-VkImageCreateInfo-pNext-06785  
  If the `pNext` chain includes a
  [VkImportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalTextureInfoEXT.html)
  structure and the image does not have a multi-planar format, then
  [VkImportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalTextureInfoEXT.html)::`plane`
  **must** be `VK_IMAGE_ASPECT_PLANE_0_BIT`

- <a href="#VUID-VkImageCreateInfo-pNext-06786"
  id="VUID-VkImageCreateInfo-pNext-06786"></a>
  VUID-VkImageCreateInfo-pNext-06786  
  If the `pNext` chain includes a
  [VkImportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalTextureInfoEXT.html)
  structure and the image has a multi-planar format with only two
  planes, then
  [VkImportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalTextureInfoEXT.html)::`plane`
  **must** not be `VK_IMAGE_ASPECT_PLANE_2_BIT`

- <a href="#VUID-VkImageCreateInfo-imageCreateFormatFeatures-09048"
  id="VUID-VkImageCreateInfo-imageCreateFormatFeatures-09048"></a>
  VUID-VkImageCreateInfo-imageCreateFormatFeatures-09048  
  If `imageCreateFormatFeatures` (as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-creation-limits"
  target="_blank" rel="noopener">Image Creation Limits</a>) does not
  contain `VK_FORMAT_FEATURE_2_HOST_IMAGE_TRANSFER_BIT_EXT`, then
  `usage` **must** not contain `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT`

- <a href="#VUID-VkImageCreateInfo-pNext-09653"
  id="VUID-VkImageCreateInfo-pNext-09653"></a>
  VUID-VkImageCreateInfo-pNext-09653  
  If the `pNext` chain contains a
  [VkImageAlignmentControlCreateInfoMESA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAlignmentControlCreateInfoMESA.html)
  structure, `tiling` **must** be `VK_IMAGE_TILING_OPTIMAL`

- <a href="#VUID-VkImageCreateInfo-pNext-09654"
  id="VUID-VkImageCreateInfo-pNext-09654"></a>
  VUID-VkImageCreateInfo-pNext-09654  
  If the `pNext` chain contains a
  [VkImageAlignmentControlCreateInfoMESA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAlignmentControlCreateInfoMESA.html)
  structure, it **must** not contain a
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
  structure

Valid Usage (Implicit)

- <a href="#VUID-VkImageCreateInfo-sType-sType"
  id="VUID-VkImageCreateInfo-sType-sType"></a>
  VUID-VkImageCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO`

- <a href="#VUID-VkImageCreateInfo-pNext-pNext"
  id="VUID-VkImageCreateInfo-pNext-pNext"></a>
  VUID-VkImageCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html),
  [VkDedicatedAllocationImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationImageCreateInfoNV.html),
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html),
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html),
  [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatQNX.html),
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html),
  [VkExternalMemoryImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfoNV.html),
  [VkImageAlignmentControlCreateInfoMESA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAlignmentControlCreateInfoMESA.html),
  [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html),
  [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html),
  [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html),
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html),
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html),
  [VkImageSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSwapchainCreateInfoKHR.html),
  [VkImportMetalIOSurfaceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalIOSurfaceInfoEXT.html),
  [VkImportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalTextureInfoEXT.html),
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html),
  [VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatInfoNV.html),
  or [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)

- <a href="#VUID-VkImageCreateInfo-sType-unique"
  id="VUID-VkImageCreateInfo-sType-unique"></a>
  VUID-VkImageCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique, with the exception of structures of type
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  or [VkImportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalTextureInfoEXT.html)

- <a href="#VUID-VkImageCreateInfo-flags-parameter"
  id="VUID-VkImageCreateInfo-flags-parameter"></a>
  VUID-VkImageCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html) values

- <a href="#VUID-VkImageCreateInfo-imageType-parameter"
  id="VUID-VkImageCreateInfo-imageType-parameter"></a>
  VUID-VkImageCreateInfo-imageType-parameter  
  `imageType` **must** be a valid [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html) value

- <a href="#VUID-VkImageCreateInfo-format-parameter"
  id="VUID-VkImageCreateInfo-format-parameter"></a>
  VUID-VkImageCreateInfo-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkImageCreateInfo-samples-parameter"
  id="VUID-VkImageCreateInfo-samples-parameter"></a>
  VUID-VkImageCreateInfo-samples-parameter  
  `samples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

- <a href="#VUID-VkImageCreateInfo-tiling-parameter"
  id="VUID-VkImageCreateInfo-tiling-parameter"></a>
  VUID-VkImageCreateInfo-tiling-parameter  
  `tiling` **must** be a valid [VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html) value

- <a href="#VUID-VkImageCreateInfo-usage-parameter"
  id="VUID-VkImageCreateInfo-usage-parameter"></a>
  VUID-VkImageCreateInfo-usage-parameter  
  `usage` **must** be a valid combination of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) values

- <a href="#VUID-VkImageCreateInfo-usage-requiredbitmask"
  id="VUID-VkImageCreateInfo-usage-requiredbitmask"></a>
  VUID-VkImageCreateInfo-usage-requiredbitmask  
  `usage` **must** not be `0`

- <a href="#VUID-VkImageCreateInfo-sharingMode-parameter"
  id="VUID-VkImageCreateInfo-sharingMode-parameter"></a>
  VUID-VkImageCreateInfo-sharingMode-parameter  
  `sharingMode` **must** be a valid [VkSharingMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html)
  value

- <a href="#VUID-VkImageCreateInfo-initialLayout-parameter"
  id="VUID-VkImageCreateInfo-initialLayout-parameter"></a>
  VUID-VkImageCreateInfo-initialLayout-parameter  
  `initialLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageMemoryRequirements.html),
[VkDeviceImageSubresourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageSubresourceInfoKHR.html),
[VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlags.html),
[VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatConstraintsInfoFUCHSIA.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html), [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html),
[VkSharingMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
